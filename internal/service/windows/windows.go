//go:build windows

package main

import (
	"log"
	"strconv"
	"time"

	"github.com/doncicuto/openuem-nats-service/internal/common"
	"github.com/doncicuto/openuem-nats-service/internal/logger"
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
	/* var natsCmd *exec.Cmd */

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	config, err := common.GenerateNatsConfig()
	if err != nil {
		return
	}

	log.Println("[INFO]: launching NATS server")

	/* exePath := common.GetNATSBinPath() */
	cfgPath := common.GetNATSConfigPath()

	flagOpts := server.Options{}

	if config.ClusterPort != "" && config.ClusterName != "" && config.OtherServers != "" {
		flagOpts.Cluster.Name = config.ClusterName
		flagOpts.Cluster.Host = config.ServerName
		flagOpts.Cluster.Port, err = strconv.Atoi(config.ClusterPort)
		if err != nil {
			log.Println("[FATAL]: could not set NATS cluster port")
			return
		}
		// TODO - Add routes to other servers using config.OtherServers
		// natsCmd = exec.Command(exePath, "--cluster_name", config.ClusterName, "-cluster", "nats://"+config.ServerName+":"+config.ClusterPort, "-routes", config.OtherServers, "-c", cfgPath)
	}

	fileOpts, err := server.ProcessConfigFile(cfgPath)
	if err != nil {
		log.Println("[FATAL]: could not parse NATS config file")
		return
	}

	opts := server.MergeOptions(fileOpts, &flagOpts)
	ns, err := server.NewServer(opts)
	if err != nil {
		log.Fatalf("server init: %v", err)
	}

	go ns.Start()
	log.Println("[INFO]: NATS embedded server has been started")

	/* go func() {
		if err := natsCmd.Start(); err != nil {
			log.Printf("[ERROR]: could not start nats command: %v", err)
			return
		}

		if err := natsCmd.Wait(); err != nil {
			log.Printf("[ERROR]: could not wait for nats command to finish: %v", err)
			return
		}

		log.Println("[INFO]: NATS server is running")
	}() */

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
