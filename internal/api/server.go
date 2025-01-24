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
	"salesforge-api/internal/service"
)

func NewServer(
	conf config.ServerConfig,
	sequenceService service.SequenceService,
	db *sql.DB,
	l *zap.Logger,
) *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.AppServerPort),
		Handler: handlers(chi.NewRouter(), sequenceService, db),
	}

	return server
}

func handlers(
	r *chi.Mux,
	sequenceService service.SequenceService,
	db *sql.DB,
) *chi.Mux {
	sequenceHandler := sequence.NewSequenceHandler(sequenceService)
	healthCheckHandler := healthcheck.NewHealthCheckHandler(db)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/sequence", sequenceHandler.AddSequence)
		r.Put("/sequence", sequenceHandler.UpdateSequence)
		r.Put("/step", sequenceHandler.UpdateStep)
		r.Delete("/step", sequenceHandler.DeleteStep)
	})

	r.Get("/health", healthCheckHandler.ServeHTTP)

	return r
}
