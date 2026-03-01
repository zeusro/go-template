package observability

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// HTTPRequestDuration tracks HTTP request duration
	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
	)

	// HTTPRequestTotal tracks total HTTP requests
	HTTPRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// HTTPRequestInFlight tracks in-flight HTTP requests
	HTTPRequestInFlight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Number of HTTP requests currently being processed",
		},
	)
)

func init() {
	prometheus.MustRegister(HTTPRequestDuration)
	prometheus.MustRegister(HTTPRequestTotal)
	prometheus.MustRegister(HTTPRequestInFlight)
}

// PrometheusMiddleware returns a Gin middleware for Prometheus metrics
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		HTTPRequestInFlight.Inc()
		defer HTTPRequestInFlight.Dec()

		start := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
			HTTPRequestDuration.WithLabelValues(
				c.Request.Method,
				c.FullPath(),
				strconv.Itoa(c.Writer.Status()),
			).Observe(v)
		}))

		c.Next()

		HTTPRequestTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			strconv.Itoa(c.Writer.Status()),
		).Inc()

		start.ObserveDuration()
	}
}

// PrometheusHandler returns a handler for Prometheus metrics endpoint
func PrometheusHandler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}
