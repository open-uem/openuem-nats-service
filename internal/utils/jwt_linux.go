//go:build linux

package openuem_utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func GetJWTKey() (string, error) {
	// Open ini file
	configFile := GetConfigFile()
	cfg, err := ini.Load(configFile)
	if err != nil {
		return "", err
	}

	key, err := cfg.Section("JWT").GetKey("key")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresDatabase from INI")
	}

	return key.String(), nil
}
