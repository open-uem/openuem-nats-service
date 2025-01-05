//go:build windows

package common

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/open-uem/utils"
)

func GetTemplatePath() string {
	wd, err := utils.GetWd()
	if err != nil {
		log.Fatalf("[FATAL]: could not get working directory")
	}
	return filepath.Join(wd, "config", "nats.tmpl")
}

func GetNATSConfigPath() string {
	wd, err := utils.GetWd()
	if err != nil {
		log.Fatalf("[FATAL]: could not get working directory")
	}
	return filepath.Join(wd, "config", "nats.conf")
}

func GetJetstreamPath() string {
	wd, err := utils.GetWd()
	if err != nil {
		log.Fatalf("[FATAL]: could not get working directory")
	}
	return strings.ReplaceAll(filepath.Join(wd, "bin", "nats"), "\\", "\\\\")
}

func GetNATSBinPath() string {
	cwd, err := utils.GetWd()
	if err != nil {
		log.Fatalf("[FATAL]: could not get working directory")
	}
	return filepath.Join(cwd, "bin", "nats", "nats-server.exe")
}
