package workers

import (
	"log_cache/models"

	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"

	"log_cache/config"
	"log_cache/workers/api"
)

const APIServer = "api_server"

func InitChief(logger *logrus.Entry, cfg *config.Cfg) uwe.Chief {
	logger = logger.WithField("app_layer", "workers")

	chief := uwe.NewChief()
	chief.UseDefaultRecover()
	chief.EnableServiceSocket(config.AppInfo())
	chief.SetEventHandler(uwe.LogrusEventHandler(logger))

	eventBus := make(chan models.Event, 16)
	chief.AddWorker(APIServer,
		api.GetServer(cfg, logger.WithField("worker", APIServer), eventBus))

	return chief
}
