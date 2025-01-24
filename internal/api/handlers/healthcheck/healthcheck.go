package healthcheck

import (
	"database/sql"
	"go.uber.org/zap"
	"net/http"
)

type HealthCheckHandler struct {
	DB     *sql.DB
	logger *zap.Logger
}

func NewHealthCheckHandler(db *sql.DB, logger *zap.Logger) *HealthCheckHandler {
	return &HealthCheckHandler{DB: db, logger: logger}
}

func (h *HealthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Healthcheck request received")
	if err := h.DB.Ping(); err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
