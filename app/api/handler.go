package api

import (
	"log_cache/config"

	"github.com/sirupsen/logrus"
)

// Handler contains realization of http handlers
type Handler struct {
	log *logrus.Entry
}

func NewHandler(cfg *config.Cfg, entry *logrus.Entry) *Handler {
	return &Handler{
		log: entry.WithField("app_layer", "api.Handler"),
	}
}
