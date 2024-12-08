//go:build windows

package logger

import (
	"log"
	"os"
	"path/filepath"

	"github.com/doncicuto/openuem_utils"
)

type OpenUEMLogger struct {
	LogFile *os.File
}

func New() *OpenUEMLogger {
	logger := OpenUEMLogger{}

	// Get executable path to store logs
	cwd, err := openuem_utils.GetWd()
	if err != nil {
		log.Fatalf("could not get executable info: %v", err)
	}

	logPath := filepath.Join(cwd, "logs", "openuem-nats-service.txt")
	logger.LogFile, err = os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("could not create log file: %v", err)
	}

	log.SetOutput(logger.LogFile)
	log.SetPrefix("openuem-nats-service: ")
	log.SetFlags(log.Ldate | log.Ltime)

	return &logger
}

func (l *OpenUEMLogger) Close() {
	l.LogFile.Close()
}
