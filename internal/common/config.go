package common

import (
	"log"
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/doncicuto/openuem_utils"
	"gopkg.in/ini.v1"
)

type NATSConfig struct {
	Org            string
	Street         string
	Locality       string
	Province       string
	Country        string
	PostalCode     string
	NATSCertPath   string
	NATSKeyPath    string
	NATSCAPath     string
	Port           string
	JetstreamStore string
	ClusterName    string
	ClusterPort    string
	OtherServers   string
	ServerName     string
	Debug          bool
}

func GenerateNatsConfig() (*NATSConfig, error) {
	var err error
	data := NATSConfig{}

	// Open ini file
	configFile := openuem_utils.GetConfigFile()

	cfg, err := ini.Load(configFile)
	if err != nil {
		return nil, err
	}

	key, err := cfg.Section("Certificates").GetKey("OrgName")
	if err != nil {
		log.Printf("[ERROR]: could not get org name, %v", err)
		return nil, err
	}
	data.Org = key.String()

	key, err = cfg.Section("Certificates").GetKey("OrgAddress")
	if err != nil {
		log.Println("[ERROR]: could not get org address")
		return nil, err
	}
	data.Street = key.String()

	key, err = cfg.Section("Certificates").GetKey("OrgCountry")
	if err != nil {
		log.Println("[ERROR]: could not get org country")
		return nil, err
	}
	data.Country = key.String()

	key, err = cfg.Section("Certificates").GetKey("OrgLocality")
	if err != nil {
		log.Println("[ERROR]: could not get org locality")
		return nil, err
	}
	data.Locality = key.String()

	key, err = cfg.Section("Certificates").GetKey("OrgProvince")
	if err != nil {
		log.Println("[ERROR]: could not get org province")
		return nil, err
	}
	data.Province = key.String()

	key, err = cfg.Section("NATS").GetKey("NATSServer")
	if err != nil {
		log.Println("[ERROR]: could not get NATS server")
		return nil, err
	}
	data.ServerName = key.String()

	key, err = cfg.Section("NATS").GetKey("NATSPort")
	if err != nil {
		log.Println("[ERROR]: could not get NATS port")
		return nil, err
	}
	data.Port = key.String()

	data.JetstreamStore = GetJetstreamPath()

	key, err = cfg.Section("Certificates").GetKey("CACert")
	if err != nil {
		log.Println("[ERROR]: could not get CA certificate path")
		return nil, err
	}
	data.NATSCAPath = key.String()
	if runtime.GOOS == "windows" {
		data.NATSCAPath = strings.ReplaceAll(data.NATSCAPath, "\\", "\\\\")
	}

	key, err = cfg.Section("Certificates").GetKey("NATSCert")
	if err != nil {
		log.Println("[ERROR]: could not get NATS certificate path")
		return nil, err
	}
	data.NATSCertPath = key.String()
	if runtime.GOOS == "windows" {
		data.NATSCertPath = strings.ReplaceAll(data.NATSCertPath, "\\", "\\\\")
	}

	key, err = cfg.Section("Certificates").GetKey("NATSKey")
	if err != nil {
		log.Println("[ERROR]: could not get NATS private key path")
		return nil, err
	}
	data.NATSKeyPath = key.String()
	if runtime.GOOS == "windows" {
		data.NATSKeyPath = strings.ReplaceAll(data.NATSKeyPath, "\\", "\\\\")
	}

	tmplPath := GetTemplatePath()
	tmpl, err := template.New("nats.tmpl").ParseFiles(tmplPath)
	if err != nil {
		log.Printf("[ERROR]: could not open the template: %v", err)
		return nil, err
	}

	cfgPath := GetNATSConfigPath()
	file, err := os.Create(cfgPath)
	if err != nil {
		log.Printf("[ERROR]:could not create the config file: %v", err)
		return nil, err
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		log.Printf("[ERROR]:could not generate the config file from the template: %v", err.Error())
		return nil, err
	}

	key, err = cfg.Section("NATS").GetKey("NATSClusterPort")
	if err == nil {
		data.ClusterPort = key.String()
	}

	key, err = cfg.Section("NATS").GetKey("NATSClusterName")
	if err == nil {
		data.ClusterName = key.String()
	}

	key, err = cfg.Section("NATS").GetKey("NATSOtherServers")
	if err == nil {
		data.OtherServers = key.String()
	}

	key, err = cfg.Section("NATS").GetKey("Debug")
	if err == nil {
		if key.String() == "true" {
			data.Debug = true
		}
	}

	return &data, nil
}
