package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"salesforge-api/internal/api/handlers/sequence"
	"salesforge-api/internal/config"
	"salesforge-api/internal/service"
)

func NewServer(
	conf config.ServerConfig,
	sequenceService service.SequenceService,
	l *zap.Logger,
) *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.AppServerPort),
		Handler: handlers(chi.NewRouter(), sequenceService),
	}

	return server
}

func handlers(
	r *chi.Mux,
	sequenceService service.SequenceService,
) *chi.Mux {
	sequenceHandler := sequence.NewSequenceHandler(sequenceService)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/sequence", sequenceHandler.AddSequence)
		r.Put("/sequence", sequenceHandler.UpdateSequence)
		r.Put("/step", sequenceHandler.UpdateStep)
		r.Delete("/step", sequenceHandler.DeleteStep)
	})

	return r
}
