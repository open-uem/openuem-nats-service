//go:build linux

package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"text/template"

	"github.com/doncicuto/openuem-nats-service/common"
	"github.com/doncicuto/openuem-nats-service/logger"
	"gopkg.in/ini.v1"
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

func main() {
	l := logger.New()

	err := generateNatsConfig()
	if err != nil {
		log.Fatalf("[FATAL]: could not generate NATS config")

	}

	log.Println("[INFO]: launching NATS server")

	exePath := "/opt/openuem-server/bin/nats/nats-server"
	cfgPath := "/etc/openuem-server/nats.cfg"

	natsCmd := exec.Command(exePath, "-c", cfgPath, "-m", "8433")

	go func() {
		if err := natsCmd.Start(); err != nil {
			log.Printf("[ERROR]: could not start nats command: %v", err)
			return
		}

		// Save pid to PIDFILE
		if err := os.WriteFile("PIDFILE", []byte(strconv.Itoa(os.Getpid())), 0666); err != nil {
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

	l.Close()
	stopCmd := exec.Command("/usr/bin/pkill", "-F", "./PIDFILE")
	if err := stopCmd.Start(); err != nil {
		log.Println("[ERROR]: could not kill nats-server process")
	}

	// TODO remove PIDFILE

	log.Printf("[INFO]: the OCSP responder has stopped listening\n")
}

func generateNatsConfig() error {
	var err error
	data := common.NATSConfig{}

	// Open ini file
	cfg, err := ini.Load("/etc/openuem-server/openuem.ini")
	if err != nil {
		return err
	}

	key, err := cfg.Section("Server").GetKey("org_name")
	if err != nil {
		log.Println("[ERROR]: could not get org name")
		return err
	}
	data.Org = key.String()

	key, err = cfg.Section("Server").GetKey("org_address")
	if err != nil {
		log.Println("[ERROR]: could not get org address")
		return err
	}
	data.Street = key.String()

	key, err = cfg.Section("Server").GetKey("org_country")
	if err != nil {
		log.Println("[ERROR]: could not get org country")
		return err
	}
	data.Country = key.String()

	key, err = cfg.Section("Server").GetKey("org_locality")
	if err != nil {
		log.Println("[ERROR]: could not get org locality")
		return err
	}
	data.Locality = key.String()

	key, err = cfg.Section("Server").GetKey("org_province")
	if err != nil {
		log.Println("[ERROR]: could not get org province")
		return err
	}
	data.Province = key.String()

	key, err = cfg.Section("Server").GetKey("nats_port")
	if err != nil {
		log.Println("[ERROR]: could not get NATS port")
		return err
	}
	data.Port = key.String()

	data.JetstreamStore = "/opt/openuem-server/bin/nats"

	key, err = cfg.Section("Server").GetKey("ca_cert_path")
	if err != nil {
		log.Println("[ERROR]: could not get CA certificate path")
		return err
	}
	data.NATSCAPath = key.String()

	key, err = cfg.Section("Server").GetKey("nats_cert_path")
	if err != nil {
		log.Println("[ERROR]: could not get NATS certificate path")
		return err
	}
	data.NATSCertPath = key.String()

	key, err = cfg.Section("Server").GetKey("nats_key_path")
	if err != nil {
		log.Println("[ERROR]: could not get NATS private key path")
		return err
	}
	data.NATSKeyPath = key.String()

	tmplPath := "/opt/openuem-server/bin/nats/nats.tmpl"
	tmpl, err := template.New("nats.tmpl").ParseFiles(tmplPath)
	if err != nil {
		log.Printf("[ERROR]: could not open the template: %v", err)
		return err
	}

	cfgPath := "/etc/openuem-server/nats.cfg"
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
