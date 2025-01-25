package api

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"salesforge-api/internal/api/handlers/healthcheck"
	"salesforge-api/internal/api/handlers/sequence"
	"salesforge-api/internal/config"
	"salesforge-api/internal/middleware"
	"salesforge-api/internal/monitoring"
	"salesforge-api/internal/service"
	"time"
)

func NewServer(
	conf config.ServerConfig,
	sequenceService service.SequenceService,
	l *zap.Logger,
) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.LoggingMiddleware(l))
	r.Use(middleware.ErrorHandlingMiddleware(l))

	if conf.JWTAuthentication {
		r.Use(middleware.Authenticate)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.AppServerPort),
		Handler: handlers(r, sequenceService, l),
	}

	return server
}

func NewHealthCheckServer(
	conf config.ServerConfig,
	db *sql.DB,
	l *zap.Logger,
) *http.Server {
	r := chi.NewRouter()
	healthCheckHandler := healthcheck.NewHealthCheckHandler(db, l)

	r.Get("/health", healthCheckHandler.ServeHTTP)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.HealthcheckPort),
		Handler: r,
	}

	return server
}

func handlers(
	r *chi.Mux,
	sequenceService service.SequenceService,
	l *zap.Logger,
) *chi.Mux {
	sequenceHandler := sequence.NewSequenceHandler(sequenceService, l)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/sequence", func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			sequenceHandler.AddSequence(w, r)
			duration := time.Since(start).Seconds()
			monitoring.RecordMetrics("/v1/sequence", duration)
		})
		r.Put("/sequence", func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			sequenceHandler.UpdateSequence(w, r)
			duration := time.Since(start).Seconds()
			monitoring.RecordMetrics("/v1/sequence", duration)
		})
		r.Put("/step", func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			sequenceHandler.UpdateStep(w, r)
			duration := time.Since(start).Seconds()
			monitoring.RecordMetrics("/v1/step", duration)
		})
		r.Delete("/step", func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			sequenceHandler.DeleteStep(w, r)
			duration := time.Since(start).Seconds()
			monitoring.RecordMetrics("/v1/step", duration)
		})
	})

	r.Get("/metrics", http.HandlerFunc(monitoring.MetricsHandler().ServeHTTP))

	return r
}
