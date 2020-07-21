package app

import (
	"log_cache/app/api"
	"log_cache/config"

	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"
)

const (
	APIServer        = "api_server"
	MonitoringServer = "monitoring_server"
)

func InitChief(logger *logrus.Entry, cfg *config.Cfg) uwe.Chief {
	logger = logger.WithField("app_layer", "workers")

	chief := uwe.NewChief()
	chief.UseDefaultRecover()
	chief.EnableServiceSocket(config.AppInfo())
	chief.SetEventHandler(uwe.LogrusEventHandler(logger))

	chief.AddWorker(APIServer, api.GetServer(cfg, logger.WithField("worker", APIServer)))
	chief.AddWorker(MonitoringServer, api.GetMonitoringServer(cfg.Monitoring))

	return chief
}
