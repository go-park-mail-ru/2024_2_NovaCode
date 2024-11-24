package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	RequestCounter    *prometheus.CounterVec
	RequestDuration   *prometheus.HistogramVec
	ErrorCounter      *prometheus.CounterVec
	ActiveConnections prometheus.Gauge
}

func New(namespace string) *Metrics {
	m := &Metrics{
		RequestCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "http_requests_total",
				Help:      "total number of http requests",
			},
			[]string{"method", "url", "status"},
		),
		RequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:      "http_request_duration_seconds",
				Help:      "histogram of response time for handler in seconds",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"method", "url"},
		),
		ErrorCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "http_errors_total",
				Help:      "total number of http errors",
			},
			[]string{"method", "url", "status"},
		),
		ActiveConnections: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "active_connections",
				Help:      "current number of active connections",
			},
		),
	}

	m.register()

	return m
}

func (m *Metrics) register() {
	prometheus.MustRegister(m.RequestCounter)
	prometheus.MustRegister(m.RequestDuration)
	prometheus.MustRegister(m.ErrorCounter)
	prometheus.MustRegister(m.ActiveConnections)
}
