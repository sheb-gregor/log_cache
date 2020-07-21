package api

import (
	"net/http"
	"time"

	"log_cache/config"
	"log_cache/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/uwe/v2/presets/api"
	"github.com/sirupsen/logrus"
)

func GetServer(cfg *config.Cfg, logger *logrus.Entry, bus chan<- models.Event) *api.Server {
	return api.NewServer(cfg.API, getRouter(cfg, logger, bus))
}

func getRouter(cfg *config.Cfg, logger *logrus.Entry, bus chan<- models.Event) http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
		log.NewRequestLogger(logger.Logger),
	)

	if cfg.API.EnableCORS {
		r.Use(getCORS().Handler)
	}

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	if cfg.API.ApiRequestTimeout > 0 {
		t := time.Duration(cfg.API.ApiRequestTimeout)
		r.Use(middleware.Timeout(t * time.Second))
	}

	// h := handler.NewHandler(cfg, logger, bus)

	r.Route("/", func(r chi.Router) {
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			render.Success(w, config.AppInfo())
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.ResultNotFound.Render(w)
	})

	return r
}

func getCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "jwt", "X-UID"},
		ExposedHeaders:   []string{"Link", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
