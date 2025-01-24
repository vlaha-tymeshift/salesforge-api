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
	"salesforge-api/internal/service"
)

func NewServer(
	conf config.ServerConfig,
	sequenceService service.SequenceService,
	db *sql.DB,
	l *zap.Logger,
) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.LoggingMiddleware(l))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.AppServerPort),
		Handler: handlers(r, sequenceService, db, l),
	}

	return server
}

func handlers(
	r *chi.Mux,
	sequenceService service.SequenceService,
	db *sql.DB,
	l *zap.Logger,
) *chi.Mux {
	sequenceHandler := sequence.NewSequenceHandler(sequenceService, l)
	healthCheckHandler := healthcheck.NewHealthCheckHandler(db, l)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/sequence", sequenceHandler.AddSequence)
		r.Put("/sequence", sequenceHandler.UpdateSequence)
		r.Put("/step", sequenceHandler.UpdateStep)
		r.Delete("/step", sequenceHandler.DeleteStep)
	})

	r.Get("/health", healthCheckHandler.ServeHTTP)

	return r
}
