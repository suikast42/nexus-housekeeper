package config

import (
	"fmt"
)

type NexusServer struct {
	Scheme     string         `koanf:"scheme"`
	Host       string         `koanf:"host"`
	Port       string         `koanf:"port"`
	Username   string         `koanf:"username"`
	Password   string         `koanf:"password"`
	KeepImages map[string]int `koanf:"keepImages"`
}

func (n NexusServer) Url() string {
	return fmt.Sprintf("%s://%s:%s/service/rest", n.Scheme, n.Host, n.Port)
}

type AppConfig struct {
	LogLevel    string      `koanf:"loglevel"`
	NexusServer NexusServer `koanf:"nexusServer"`
}

func (a NexusServer) String() string {
	return fmt.Sprintf("NexusServer.Scheme %s\nHost: %s\nNexusServer.Port: %s\nNexusServer.userName: %s",
		a.Scheme,
		a.Host,
		a.Port,
		a.Username)
}

// Stringer interface
func (a AppConfig) String() string {
	return fmt.Sprintf("LogLevel %s\n%s", a.LogLevel, a.NexusServer)
}
