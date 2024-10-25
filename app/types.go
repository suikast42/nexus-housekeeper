package app

import (
	"context"
	"crypto/tls"
	b64 "encoding/base64"
	"github.com/rs/zerolog"
	"github.com/suikast42/nexus-housekeeper/client/nexus3"
	"github.com/suikast42/nexus-housekeeper/config"
	"github.com/suikast42/nexus-housekeeper/logging"
	"net/http"
	"sync"
)

type App struct {
	log         *zerolog.Logger
	cfg         config.AppConfig
	ctx         context.Context
	nexusClient *nexus3.APIClient
	version     string
}

func (a *App) Logger() *zerolog.Logger {
	return a.log
}

func (a *App) Config() config.AppConfig {
	return a.cfg
}

func (a *App) Context() context.Context {
	return a.ctx
}

func (a *App) Version() string {
	return a.version
}

func (a *App) ClientNexus() *nexus3.APIClient {
	return a.nexusClient
}

var singleInstance *App
var lock = &sync.Mutex{}

func AppInstance() *App {
	lock.Lock()
	defer lock.Unlock()
	if singleInstance == nil {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		withContext := logging.LoggerWithContext(context.Background())
		background := context.Background()
		config := config.LoadConfig(withContext)
		client := nexus3.NewAPIClient(&nexus3.Configuration{
			BasePath: config.NexusServer.Url(),
			Host:     config.NexusServer.Host,
			Scheme:   config.NexusServer.Scheme,
			//DefaultHeader: map[string]string{"X-Grafana-Org-Id": "1", "Authorization": "Bearer " + instance.cfg.GrafanaServer.ApiKey},
			DefaultHeader: map[string]string{"Authorization": "Basic " + b64.StdEncoding.EncodeToString([]byte(config.NexusServer.Username+":"+config.NexusServer.Password))},
			UserAgent:     "nexus-housekeeper",
			HTTPClient: &http.Client{
				Transport: tr,
			},
		})

		singleInstance = &App{
			ctx:         background,
			log:         withContext,
			cfg:         config,
			nexusClient: client,
		}
	}
	return singleInstance
}
