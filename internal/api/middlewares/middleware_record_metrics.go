package middlewares

import (
	"net/http"
	"time"

	"github.com/alex-arraga/rss_project/internal/logger"
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			duration := time.Since(start)
			logger.RecordHTTPRequests(r.Method, r.URL.Path, duration)
		}()
		next.ServeHTTP(w, r)
	})
}
