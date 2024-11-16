//go:build linux

package logger

import (
	"log"
	"os"
)

type OpenUEMLogger struct {
	LogFile *os.File
}

func New() *OpenUEMLogger {
	var err error
	logger := OpenUEMLogger{}

	os.MkdirAll("/var/log/openuem-server", 0755)
	logPath := "/var/log/openuem-server/openuem-nats-service"
	logger.LogFile, err = os.Create(logPath)
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
