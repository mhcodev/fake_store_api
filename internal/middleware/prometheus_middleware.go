package middleware

import (
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

func RegisterPrometheusMetrics() {
	getAPILatencyMetric()
	getAPIRequestMetric()
	getAPIFrequencyMetric()
}

func RecordRequestLatency(c *fiber.Ctx) error {
	start := time.Now()
	next := c.Next()

	elapsed := time.Since(start).Seconds()

	if isValidPath(c.Path()) {
		apiLatency.WithLabelValues(c.Method(), normalizePath(c.Path())).Observe(elapsed)
	}

	return next
}

func RecordRequestCount(c *fiber.Ctx) error {
	next := c.Next()

	if isValidPath(c.Path()) {
		apiRequests.WithLabelValues(c.Method(), normalizePath(c.Path())).Inc()
	}

	return next
}

func RecordRequestFrequency(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start).Seconds()

	if isValidPath(c.Path()) {
		apiFrequency.WithLabelValues(c.Method(), normalizePath(c.Path())).Observe(duration)
		apiRequests.WithLabelValues(c.Method(), normalizePath(c.Path())).Inc()
	}

	return err
}

var (
	apiLatency    *prometheus.SummaryVec
	apiRequests   *prometheus.CounterVec
	apiFrequency  *prometheus.HistogramVec
	latencyOnce   sync.Once
	requestsOnce  sync.Once
	frequencyOnce sync.Once
)

func getAPILatencyMetric() {
	latencyOnce.Do(func() {
		apiLatency = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Namespace:  "api",
				Name:       "latency_seconds",
				Help:       "Latency distributions.",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
			},
			[]string{"method", "path"},
		)
		prometheus.MustRegister(apiLatency)
	})
}

func getAPIRequestMetric() {
	requestsOnce.Do(func() {
		apiRequests = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "api",
				Name:      "requests_total",
				Help:      "Total API requests by endpoint",
			},
			[]string{"method", "path"},
		)
		prometheus.MustRegister(apiRequests)
	})

}

func getAPIFrequencyMetric() {
	frequencyOnce.Do(func() {
		apiFrequency = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "api",
				Name:      "requests_frequency",
				Help:      "API requests Frequency by endpoint",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"method", "path"},
		)
		prometheus.MustRegister(apiFrequency)
	})

}

func isValidPath(path string) bool {
	return !strings.Contains(path, "metric")
}

func normalizePath(path string) string {
	if strings.HasPrefix(path, "/api/v1/user/") {
		return "/api/v1/user/:id"
	}
	if strings.HasPrefix(path, "/api/v1/category/") {
		return "/api/v1/category/:id"
	}
	if strings.HasPrefix(path, "/api/v1/product/") {
		return "/api/v1/product/:id"
	}
	if strings.HasPrefix(path, "/api/v1/file/") {
		return "/api/v1/file/:id"
	}
	return path
}
