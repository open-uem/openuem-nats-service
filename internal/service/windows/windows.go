//go:build windows

package main

import (
	"log"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/doncicuto/openuem-nats-service/internal/common"
	"github.com/doncicuto/openuem-nats-service/internal/logger"
	"github.com/doncicuto/openuem_utils"
	"github.com/nats-io/nats-server/v2/server"
	"golang.org/x/sys/windows/svc"
)

type OpenUEMService struct {
	Logger *logger.OpenUEMLogger
}

func New(l *logger.OpenUEMLogger) *OpenUEMService {
	return &OpenUEMService{
		Logger: l,
	}
}

func (s *OpenUEMService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	var err error
	var flagOpts *server.Options

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	config, err := common.GenerateNatsConfig()
	if err != nil {
		return
	}

	log.Println("[INFO]: launching NATS server")

	cfgPath := common.GetNATSConfigPath()
	fileOpts, err := server.ProcessConfigFile(cfgPath)
	if err != nil {
		log.Println("[FATAL]: could not parse NATS config file")
		return
	}

	cwd, _ := openuem_utils.GetWd()

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
		fileOpts.Debug = true
		fileOpts.Trace = true
		fileOpts.TraceVerbose = true
		fileOpts.LogFile = filepath.Join(cwd, "logs", "nats-cluster-log.txt")
	}

	ns, err := server.NewServer(fileOpts)
	if err != nil {
		log.Fatalf("server init: %v", err)
	}

	go ns.Start()
	log.Println("[INFO]: NATS embedded server has been started")

	// Add log
	ns.ConfigureLogger()

	// service control manager
loop:
	for {
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				log.Println("[INFO]: service has received the stop or shutdown command")

				ns.Shutdown()
				log.Println("[INFO]: NATS embedded server shutdown")
				/* kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(natsCmd.Process.Pid))
				if err := kill.Run(); err != nil {
					log.Printf("[ERROR]: could not kill NATS server process %v", err.Error())
				} */

				s.Logger.Close()
				break loop
			default:
				log.Println("[WARN]: unexpected control request")
				return true, 1
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return true, 0
}

func main() {
	l := logger.New()

	s := New(l)

	// Run service
	err := svc.Run("openuem-nats-service", s)
	if err != nil {
		log.Printf("[ERROR]: could not run service: %v", err)
	}

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
