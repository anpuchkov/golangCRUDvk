package middleware

import (
	"go.uber.org/zap"
	"net/http"
)

// LoggingMiddleware /* Используем Middleware для логгирования входящих и исходящих HTTP-запросов.
// Zap для будущего логгирования и записи логов в файл.
func LoggingMiddleware(logger *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(
			"incoming http request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("query", r.URL.RawQuery),
		)

		next.ServeHTTP(w, r)

		//logger.Info(
		//	"outgoing http request",
		//	zap.String("path", r.URL.Path),
		//	zap.String("query", r.URL.RawQuery),
		//	zap.String("status", strconv.Itoa(http.StatusOK)),
		//)
	})
}
