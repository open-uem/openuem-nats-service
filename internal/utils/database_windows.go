//go:build windows

package openuem_utils

import (
	"encoding/binary"
	"fmt"

	"github.com/danieljoos/wincred"
	"gopkg.in/ini.v1"
)

func CreatePostgresDatabaseURL() (string, error) {
	var err error

	// Open ini file
	configFile := GetConfigFile()
	cfg, err := ini.Load(configFile)
	if err != nil {
		return "", err
	}

	user, err := cfg.Section("DB").GetKey("PostgresUser")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresUser from INI")
	}

	host, err := cfg.Section("DB").GetKey("PostgresHost")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresHost from INI")
	}

	port, err := cfg.Section("DB").GetKey("PostgresPort")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresPort from INI")
	}

	database, err := cfg.Section("DB").GetKey("PostgresDatabase")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresDatabase from INI")
	}

	pass, err := wincred.GetGenericCredential(host.String() + ":" + port.String())
	if err != nil {
		return "", fmt.Errorf("could not read password from Windows Credential Manager")
	}

	decodedPass := UTF16BytesToString(pass.CredentialBlob, binary.LittleEndian)
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, decodedPass, host, port, database), nil
}
