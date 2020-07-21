package api

import (
	"net/http"
	"time"

	"log_cache/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/lancer-kit/armory/api/httpx"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/uwe/v2/presets/api"
	"github.com/sirupsen/logrus"
)

type LogEntry map[string]interface{}

const IPKey = "ip"

func NewLogsServer(cfg *config.Cfg, ipBus chan<- string, logger *logrus.Entry) *api.Server {
	mux := chi.NewRouter()

	// A good base middleware stack
	mux.Use(
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
		log.NewRequestLogger(logger.Logger),
	)

	if cfg.API.EnableCORS {
		mux.Use(getCORS().Handler)
	}

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	if cfg.API.ApiRequestTimeout > 0 {
		t := time.Duration(cfg.API.ApiRequestTimeout)
		mux.Use(middleware.Timeout(t * time.Second))
	}

	mux.Route("/", func(r chi.Router) {
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			render.Success(w, config.AppInfo())
		})

		r.Post("/logs", func(w http.ResponseWriter, r *http.Request) {
			logRequest := make(LogEntry)

			err := httpx.ParseJSONBody(r, &logRequest)
			if err != nil {
				render.BadRequest(w, err)
				return
			}

			ip, ok := logRequest[IPKey].(string)
			if !ok {
				render.BadRequest(w, "log data does not contain an IP field")
				return
			}

			ipBus <- ip
			render.ResultSuccess.Render(w)
		})

	})

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.ResultNotFound.Render(w)
	})

	return api.NewServer(cfg.API, mux)
}

func getCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
