//go:build linux

package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/open-uem/openuem-nats-service/internal/common"
	"github.com/open-uem/openuem-nats-service/internal/logger"
	"github.com/open-uem/utils"
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
	var flagOpts *server.Options

	l := logger.New()

	config, err := common.GenerateNatsConfig()
	if err != nil {
		log.Fatalf("[FATAL]: could not generate NATS config")

	}

	log.Println("[INFO]: launching NATS server")

	cfgPath := common.GetNATSConfigPath()

	fileOpts, err := server.ProcessConfigFile(cfgPath)
	if err != nil {
		log.Println("[FATAL]: could not parse NATS config file")
		return
	}

	cwd, err := utils.GetWd()
	if err != nil {
		log.Println("[FATAL]: could not get working directory")
		return
	}

	if config.ClusterPort != "" && config.ClusterName != "" && config.OtherServers != "" {
		flagOpts, err = GetFlagsOptions(config)
		if err != nil {
			log.Printf("[FATAL]: could not create Flags options, reason: %v", err)
			return
		}
		fileOpts.Cluster.Name = flagOpts.Cluster.Name
		fileOpts.Cluster.Port = flagOpts.Cluster.Port
		fileOpts.HTTPPort = 8222
		fileOpts.Routes = flagOpts.Routes
	}

	if config.Debug {
		fileOpts.Debug = true
		fileOpts.Trace = true
		fileOpts.LogFile = filepath.Join(cwd, "logs", "openuem-nats-debug.txt")
	}

	ns, err := server.NewServer(fileOpts)
	if err != nil {
		log.Fatalf("server init: %v", err)
	}

	go ns.Start()
	log.Println("[INFO]: NATS embedded server has been started")

	// Add log
	if config.Debug {
		ns.ConfigureLogger()
	}

	// Keep the connection alive
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	log.Println("[INFO]: the NATS server is ready and listening")
	<-done

	log.Println("[INFO]: service has received the stop or shutdown command")

	ns.Shutdown()
	log.Println("[INFO]: NATS embedded server shutdown")
	l.Close()
}

func GetFlagsOptions(config *common.NATSConfig) (*server.Options, error) {
	var err error

	flagOpts := server.Options{}
	flagOpts.Cluster.Name = config.ClusterName
	flagOpts.Cluster.Host = config.ServerName
	flagOpts.Cluster.Port, err = strconv.Atoi(config.ClusterPort)
	if err != nil {
		return nil, err
	}

	flagOpts.Routes = []*url.URL{}
	otherServers := strings.Split(config.OtherServers, ",")
	for _, server := range otherServers {
		u, err := url.Parse("tls://" + server)
		if err == nil {
			flagOpts.Routes = append(flagOpts.Routes, u)
		}
	}
	return &flagOpts, nil
}
