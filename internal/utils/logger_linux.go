//go:build linux

package openuem_utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type OpenUEMLogger struct {
	LogFile *os.File
}

func NewLogger(logFilename string) *OpenUEMLogger {
	var err error

	logger := OpenUEMLogger{}

	// Get executable path to store logs
	wd := "/var/log/openuem-server"

	if _, err := os.Stat(wd); os.IsNotExist(err) {
		if err := os.MkdirAll(wd, 0660); err != nil {
			log.Fatalf("[FATAL]: could not create log directory, reason: %v", err)
		}
	}

	logPath := filepath.Join(wd, logFilename)
	logger.LogFile, err = os.Create(logPath)
	if err != nil {
		log.Fatalf("could not create log file: %v", err)
	}

	logPrefix := strings.TrimSuffix(filepath.Base(logFilename), filepath.Ext(logFilename))
	log.SetOutput(logger.LogFile)
	log.SetPrefix(logPrefix + ": ")
	log.SetFlags(log.Ldate | log.Ltime)

	return &logger
}

func (l *OpenUEMLogger) Close() {
	l.LogFile.Close()
}
