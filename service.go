package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/doncicuto/openuem-nats-service/logger"
	"golang.org/x/sys/windows/registry"
	"golang.org/x/sys/windows/svc"
)

type NATSConfig struct {
	Org          string
	Street       string
	Locality     string
	Province     string
	Country      string
	PostalCode   string
	NATSCertPath string
	NATSKeyPath  string
	NATSCAPath   string
	Port         string
}

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
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	err = generateNatsConfig()
	if err != nil {
		return
	}

	// TODO may we set the port from registry key and avoid harcoding it?
	log.Println("[INFO]: launching NATS server")

	cwd, err := getWd()
	if err != nil {
		log.Println("[ERROR]: could not get working directory")
		return
	}

	exePath := filepath.Join(cwd, "bin", "nats", "nats-server.exe")
	cfgPath := filepath.Join(cwd, "config", "nats.cfg")

	natsCmd := exec.Command(exePath, "-c", cfgPath, "-m", "8433")

	go func() {
		if err := natsCmd.Start(); err != nil {
			log.Printf("[ERROR]: could not start nats command: %v", err)
			return
		}

		log.Println("[INFO]: NATS server is running")

		if err := natsCmd.Wait(); err != nil {
			log.Printf("[ERROR]: could not wait for nats command to finish: %v", err)
			return
		}
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

func generateNatsConfig() error {
	var err error
	data := NATSConfig{}

	k, err := openRegistry()
	if err != nil {
		log.Println("[ERROR]: could not open registry")
		return err
	}
	defer k.Close()

	data.Org, err = getValueFromRegistry(k, "OrgName")
	if err != nil {
		log.Println("[ERROR]: could not get org name")
		return err
	}

	data.Street, err = getValueFromRegistry(k, "OrgAddress")
	if err != nil {
		log.Println("[ERROR]: could not get org address")
		return err
	}

	data.Country, err = getValueFromRegistry(k, "OrgCountry")
	if err != nil {
		log.Println("[ERROR]: could not get org country")
		return err
	}

	data.Locality, err = getValueFromRegistry(k, "OrgLocality")
	if err != nil {
		log.Println("[ERROR]: could not get org locality")
		return err
	}

	data.Province, err = getValueFromRegistry(k, "OrgProvince")
	if err != nil {
		log.Println("[ERROR]: could not get org province")
		return err
	}

	data.Port, err = getValueFromRegistry(k, "NATSPort")
	if err != nil {
		log.Println("[ERROR]: could not get NATS Port")
		data.Port = "4433"
	}

	cwd, err := getWd()
	if err != nil {
		log.Println("[ERROR]: could not get workin directory")
		return err
	}

	data.NATSCAPath = strings.ReplaceAll(filepath.Join(cwd, "certificates/ca", "ca.cer"), "\\", "\\\\")
	data.NATSCertPath = strings.ReplaceAll(filepath.Join(cwd, "certificates/nats", "nats.cer"), "\\", "\\\\")
	data.NATSKeyPath = strings.ReplaceAll(filepath.Join(cwd, "certificates/nats", "nats.key"), "\\", "\\\\")

	tmplPath := filepath.Join(cwd, "config", "nats.tmpl")
	tmpl, err := template.New("nats.tmpl").ParseFiles(tmplPath)
	if err != nil {
		log.Printf("[ERROR]:could not open the template: %v", err)
		return err
	}

	cfgPath := filepath.Join(cwd, "config", "nats.cfg")
	file, err := os.Create(cfgPath)
	if err != nil {
		log.Printf("[ERROR]:could not create the config file: %v", err)
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		log.Printf("[ERROR]:could not generate the config file from the template: %v", err.Error())
		return err
	}

	return nil
}

func getWd() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		log.Printf("[ERROR]:could not get executable info: %v", err)
		return "", err
	}
	return filepath.Dir(ex), nil
}
func openRegistry() (registry.Key, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\OpenUEM\Server`, registry.QUERY_VALUE)
	if err != nil {
		return k, err
	}
	return k, nil
}

func getValueFromRegistry(k registry.Key, key string) (string, error) {
	s, _, err := k.GetStringValue(key)
	if err != nil {
		return "", err
	}

	return s, nil
}
