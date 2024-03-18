package middleware

import (
	"go.uber.org/zap"
	"net/http"
)

func LoggingMiddleware(logger *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(
			"incoming http request",
			zap.String("path", r.URL.Path),
			zap.String("query", r.URL.RawQuery),
		)

		next.ServeHTTP(w, r)
	})
}
