//go:build linux

package common

func GetTemplatePath() string {
	return "/opt/openuem-server/bin/nats/nats.tmpl"
}

func GetNATSConfigPath() string {
	return "/opt/openuem-server/bin/nats/nats.conf"
}

func GetJetstreamPath() string {
	return "/opt/openuem-server/bin/nats"
}

func GetNATSBinPath() string {
	return "/opt/openuem-server/bin/nats/nats-server"
}
