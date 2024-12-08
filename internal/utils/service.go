//go:build windows

package openuem_utils

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

type OpenUEMWindowsService struct {
	ServiceStart func()
	ServiceStop  func()
}

func NewOpenUEMWindowsService() *OpenUEMWindowsService {
	return &OpenUEMWindowsService{}
}

func (s *OpenUEMWindowsService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	s.ServiceStart()

	// service control manager loop
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
				s.ServiceStop()
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

func WindowsStartService(name string) error {
	// Ref: https://cs.opensource.google/go/x/sys/+/refs/tags/v0.27.0:windows/svc/example/manage.go
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer func() {
		if err := m.Disconnect(); err != nil {
			log.Printf("[ERROR]: could not disconnect from manager: %v", err)
		}
	}()
	s, err := m.OpenService(name)
	if err != nil {
		return fmt.Errorf("could not access service: %v", err)
	}
	defer s.Close()
	err = s.Start("is", "manual-started")
	if err != nil {
		return fmt.Errorf("could not start service: %v", err)
	}
	return nil
}

func WindowsSvcControl(serviceName string, c svc.Cmd, to svc.State) error {
	// Ref: https://cs.opensource.google/go/x/sys/+/refs/tags/v0.27.0:windows/svc/example/manage.go

	m, err := mgr.Connect()
	if err != nil {
		return fmt.Errorf("could not connect with Windows service manager: %v", err)
	}
	defer func() {
		if err := m.Disconnect(); err != nil {
			log.Printf("[ERROR]: could not disconnect from manager: %v", err)
		}
	}()

	s, err := m.OpenService(serviceName)
	if err != nil {
		return fmt.Errorf("could not access service: %v", err)
	}

	status, err := s.Control(c)
	if err != nil {
		return fmt.Errorf("could not send control=%d: %v", c, err)
	}

	timeout := time.Now().Add(10 * time.Second)
	for status.State != to {
		if timeout.Before(time.Now()) {
			return fmt.Errorf("timeout waiting for service to go to state=%d", to)
		}
		time.Sleep(300 * time.Millisecond)
		status, err = s.Query()
		if err != nil {
			return fmt.Errorf("could not retrieve service status: %v", err)
		}
	}
	return nil
}
