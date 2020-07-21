package handler

import (
	"log_cache/config"
	"log_cache/models"

	"github.com/sirupsen/logrus"
)

// Handler contains realization of http handlers
type Handler struct {
	log *logrus.Entry
	bus chan<- models.Event
}

func NewHandler(cfg *config.Cfg, entry *logrus.Entry, bus chan<- models.Event) *Handler {
	return &Handler{
		log: entry.WithField("app_layer", "api.Handler"),
		bus: bus,
	}

}
