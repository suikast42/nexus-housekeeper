package config

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
)

type NexusServer struct {
	Scheme   string `koanf:"scheme"`
	Host     string `koanf:"host"`
	Port     string `koanf:"port"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
}
type AppConfig struct {
	LogLevel    string      `koanf:"loglevel"`
	NexusServer NexusServer `koanf:"nexusServer"`
}

func (a *App) Logger() *zerolog.Logger {
	return a.log
}

func (a *App) Config() AppConfig {
	return a.cfg
}

func (a *App) Context() context.Context {
	return a.ctx
}

func (a *App) Version() string {
	return a.version
}

func (a NexusServer) String() string {
	return fmt.Sprintf("NexusServer.Scheme %s\nHost: %s\nNexusServer.Port: %s\nNexusServer.userName: %s", a.Scheme, a.Host, a.Port, a.Username)
}

// Stringer interface
func (a AppConfig) String() string {
	return fmt.Sprintf("LogLevel %s\n%s", a.LogLevel, a.NexusServer)
}

type App struct {
	log     *zerolog.Logger
	cfg     AppConfig
	ctx     context.Context
	version string
}
