package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	Microservice string

	RequestCounter  *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
	ErrorCounter    *prometheus.CounterVec
}

func New(namespace, microservce string) *Metrics {
	m := &Metrics{
		Microservice: microservce,
		RequestCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "http_requests_total",
				Help:      "total number of http requests",
			},
			[]string{"method", "url", "status", "microservice"},
		),
		RequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:      "http_request_duration_seconds",
				Help:      "histogram of response time for handler in seconds",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"method", "url", "microservice"},
		),
		ErrorCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "http_errors_total",
				Help:      "total number of http errors",
			},
			[]string{"method", "url", "status", "microservice"},
		),
	}

	m.register()

	return m
}

func (m *Metrics) register() {
	prometheus.MustRegister(m.RequestCounter)
	prometheus.MustRegister(m.RequestDuration)
	prometheus.MustRegister(m.ErrorCounter)
}
