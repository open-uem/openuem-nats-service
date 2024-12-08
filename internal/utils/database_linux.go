//go:build linux

package openuem_utils

import (
	"gopkg.in/ini.v1"
)

func CreatePostgresDatabaseURL() (string, error) {
	var err error

	// Open ini file
	cfg, err := ini.Load("/etc/openuem-server/openuem.ini")
	if err != nil {
		return "", err
	}

	key, err := cfg.Section("Server").GetKey("db_url")
	if err != nil {
		return "", err
	}

	return key.String(), nil
}
