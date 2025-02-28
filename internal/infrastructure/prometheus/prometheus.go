// internal/repository/prometheus_repository.go
package prometheus

import (
	"context"

	"github.com/probuborka/NutriAI/internal/entity"
	"github.com/prometheus/client_golang/prometheus"
)

type prometheusRepository struct {
	counter   *prometheus.CounterVec
	histogram *prometheus.HistogramVec
}

func NewPrometheus() *prometheusRepository {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	prometheus.MustRegister(counter, histogram)
	return &prometheusRepository{counter: counter, histogram: histogram}
}

func (r *prometheusRepository) Save(ctx context.Context, metric entity.Metric) error {
	switch metric.Type {
	case entity.MetricTypeCounter:
		r.counter.With(metric.Labels).Add(metric.Value)
	case entity.MetricTypeHistogram:
		r.histogram.With(metric.Labels).Observe(metric.Value)
	}
	return nil
}
