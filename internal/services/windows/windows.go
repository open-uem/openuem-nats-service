//go:build windows

package main

import (
	"log"
	"os/exec"
	"strconv"
	"time"

	"github.com/doncicuto/openuem-nats-service/internal/common"
	"github.com/doncicuto/openuem-nats-service/internal/logger"
	"github.com/doncicuto/openuem-nats-service/internal/models"
	"github.com/doncicuto/openuem_ent/component"
	"github.com/doncicuto/openuem_utils"
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
	var natsCmd *exec.Cmd

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	dbUrl, err := openuem_utils.CreatePostgresDatabaseURL()
	if err != nil {
		log.Println("[ERROR]: could not get database url")
		return
	}

	model, err := models.New(dbUrl)
	if err != nil {
		log.Println("[ERROR]: could not connect with database")
		return
	}
	log.Println("[INFO]: connected to database")

	// Save component version
	if err := model.SetComponent(component.ComponentNats, common.VERSION, common.CHANNEL); err != nil {
		log.Fatalf("[ERROR]: could not save component information")
	}
	log.Println("[INFO]: component information saved")

	config, err := common.GenerateNatsConfig()
	if err != nil {
		return
	}

	// TODO may we set the port from registry key and avoid harcoding it?
	log.Println("[INFO]: launching NATS server")

	exePath := common.GetNATSBinPath()
	cfgPath := common.GetNATSConfigPath()

	if config.ClusterPort != "" && config.ClusterName != "" && config.OtherServers != "" {
		natsCmd = exec.Command(exePath, "--cluster_name", config.ClusterName, "-cluster", "nats://"+config.ServerName+":"+config.ClusterPort, "-routes", config.OtherServers, "-c", cfgPath)
	} else {
		natsCmd = exec.Command(exePath, "-c", cfgPath)
	}

	go func() {
		if err := natsCmd.Start(); err != nil {
			log.Printf("[ERROR]: could not start nats command: %v", err)
			return
		}

		if err := natsCmd.Wait(); err != nil {
			log.Printf("[ERROR]: could not wait for nats command to finish: %v", err)
			return
		}

		log.Println("[INFO]: NATS server is running")
	}()

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

				kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(natsCmd.Process.Pid))
				if err := kill.Run(); err != nil {
					log.Printf("[ERROR]: could not kill NATS server process %v", err.Error())
				}

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
