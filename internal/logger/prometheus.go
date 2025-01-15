package logger

import (
	"errors"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

var (
	myapp = "RSS_Project"

	// Métrica para contar operaciones procesadas
	opsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: myapp + "_ops_total",
		Help: "The total number of processed events",
	})

	// Métrica para la cantidad de solicitudes HTTP
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: myapp + "_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	// Métrica para medir la duración de las solicitudes HTTP
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    myapp + "_http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// Métrica para registrar errores
	errorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: myapp + "_error_count_total",
			Help: "Total number of errors logged",
		},
		[]string{"type"},
	)
)

func init() {
	// Registrar las métricas personalizadas
	prometheus.MustRegister(opsProcessed)
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(errorCount)
}

// Incrementa las métricas de solicitudes HTTP
func RecordHTTPRequests(method, endpoint string, duration time.Duration) {
	httpRequestsTotal.WithLabelValues(method, endpoint).Inc()
	httpRequestDuration.WithLabelValues(method, endpoint).Observe(duration.Seconds())
}

// Incrementa las métricas de errores
func RecordError(errorType string) {
	errorCount.WithLabelValues(errorType).Inc()
}

// Inicia el servidor de métricas de Prometheus
func StartPrometheus() error {
	// Escuchar servidor en /metrics
	log.Info().Msg("Prometheus metrics server starting on port 2112")

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		return errors.New("prometheus server failed")
	}

	return nil
}
