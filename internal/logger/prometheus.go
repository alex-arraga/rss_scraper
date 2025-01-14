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
	// Métrica para contar operaciones procesadas
	opsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	// Métrica para la cantidad de solicitudes HTTP
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	// Métrica para medir la duración de las solicitudes HTTP
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "myapp_http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// Métrica para registrar errores
	errorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_error_count_total",
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

// Recorda métricas procesadas cada 2 segundos como ejemplo
func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

// Incrementa las métricas de solicitudes HTTP
func RecordHTTPRequest(method, endpoint string, duration time.Duration) {
	httpRequestsTotal.WithLabelValues(method, endpoint).Inc()
	httpRequestDuration.WithLabelValues(method, endpoint).Observe(duration.Seconds())
}

// Incrementa las métricas de errores
func RecordError(errorType string) {
	errorCount.WithLabelValues(errorType).Inc()
}

// Inicia el servidor de métricas de Prometheus
func StartPrometheus() error {
	log.Info().Msg("Prometheus recording metrics...")
	recordMetrics()

	// Escuchar servidor en /metrics
	log.Info().Msg("Prometheus metrics server starting on port 2112")

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		return errors.New("prometheus server failed")
	}

	return nil
}
