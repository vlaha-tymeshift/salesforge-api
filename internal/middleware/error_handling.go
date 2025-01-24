package middleware

import (
	"errors"
	"go.uber.org/zap"
	"net/http"
	sfErr "salesforge-api/internal/errors"
)

func ErrorHandlingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					var appErr *sfErr.AppError
					if err, ok := rec.(error); ok && errors.As(err, &appErr) {
						logger.Error("handled error", zap.Error(appErr))
						http.Error(w, appErr.Message, appErr.Code)
					} else {
						logger.Error("unhandled error", zap.Any("error", rec))
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					}
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
