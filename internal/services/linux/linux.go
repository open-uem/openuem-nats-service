//go:build linux

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/doncicuto/openuem-nats-service/internal/common"
	"github.com/doncicuto/openuem-nats-service/internal/logger"
)

type OpenUEMService struct {
	Logger *logger.OpenUEMLogger
}

func New(l *logger.OpenUEMLogger) *OpenUEMService {
	return &OpenUEMService{
		Logger: l,
	}
}

func main() {
	var natsCmd *exec.Cmd

	l := logger.New()

	config, err := common.GenerateNatsConfig()
	if err != nil {
		log.Fatalf("[FATAL]: could not generate NATS config")

	}

	log.Println("[INFO]: launching NATS server")

	exePath := common.GetNATSBinPath()
	cfgPath := common.GetNATSConfigPath()

	if config.ClusterPort != "" && config.ClusterName != "" && config.OtherServers != "" {
		natsCmd = exec.Command(exePath, "--cluster_name", config.ClusterName, "-cluster", "nats://"+config.ServerName+":"+config.ClusterPort, "-routes", config.OtherServers, "-c", cfgPath)
	} else {
		natsCmd = exec.Command(exePath, "-c", cfgPath)
	}

	// Save pid to PIDFILE
	wd, err := getWd()
	if err != nil {
		log.Fatal("[FATAL]: could not get working directory")
	}

	pidFile := filepath.Join(wd, "NATSPIDFILE")

	go func() {
		if err := natsCmd.Start(); err != nil {
			log.Printf("[ERROR]: could not start nats command: %v", err)
			return
		}

		if err := os.WriteFile(pidFile, []byte(fmt.Sprintf("%d\n", os.Getpid())), 0666); err != nil {
			log.Fatal("[FATAL]: could not create pid file")
		}

		log.Println("[INFO]: NATS server is running")

		if err := natsCmd.Wait(); err != nil {
			log.Printf("[ERROR]: could not wait for nats command to finish: %v", err)
			return
		}
	}()

	// Keep the connection alive
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	log.Println("[INFO]: the NATS server is ready and listening")
	<-done

	stopCmd := exec.Command("/usr/bin/pkill", "-F", pidFile)
	if err := stopCmd.Start(); err != nil {
		log.Println("[ERROR]: could not kill nats-server process")
	}

	// TODO remove PIDFILE
	if err := os.Remove(pidFile); err != nil {
		log.Println("[ERROR]: could not remove PIDFILE")
	}
	log.Printf("[INFO]: the NATS responder has stopped listening\n")
	l.Close()
}

func getWd() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		log.Printf("[ERROR]:could not get executable info: %v", err)
		return "", err
	}
	return filepath.Dir(ex), nil
}
