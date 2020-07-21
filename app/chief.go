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
	IPCollector      = "ip_collector"
)

func InitChief(logger *logrus.Entry, cfg *config.Cfg) uwe.Chief {
	logger = logger.WithField("app_layer", "workers")

	chief := uwe.NewChief()
	chief.UseDefaultRecover()
	chief.EnableServiceSocket(config.AppInfo())
	chief.SetEventHandler(uwe.LogrusEventHandler(logger))

	ipBus := make(chan string, 8)

	chief.AddWorkers(map[uwe.WorkerName]uwe.Worker{
		APIServer:        api.NewLogsServer(cfg, ipBus, logger.WithField("worker", APIServer)),
		MonitoringServer: api.NewMonitoringServer(cfg.Monitoring),
		IPCollector:      NewCollector(ipBus, logger.WithField("worker", IPCollector)),
	})

	return chief
}
