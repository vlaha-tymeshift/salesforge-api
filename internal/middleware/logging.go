// internal/middleware/logging.go
package middleware

import (
	"bytes"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var bodyBytes []byte
			if r.Body != nil {
				bodyBytes, _ = io.ReadAll(r.Body)
			}
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			logger.Info("request received",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
				zap.String("body", string(bodyBytes)),
			)

			next.ServeHTTP(w, r)
		})
	}
}
