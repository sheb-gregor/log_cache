package api

import (
	"net/http"

	"log_cache/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/uwe/v2/presets/api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewMonitoringServer(cfg api.Config) *api.Server {
	mux := chi.NewRouter()

	// A good base middleware stack
	mux.Use(
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
	)

	mux.Handle("/metrics", promhttp.Handler())

	mux.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		render.Success(w, config.AppInfo())
	})

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.ResultNotFound.Render(w)
	})

	return api.NewServer(cfg, mux)
}
