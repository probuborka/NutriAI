package http

import (
	"context"
	"net/http"
	"time"

	"github.com/probuborka/NutriAI/internal/entity"
)

type metric interface {
	RecordMetric(ctx context.Context, metric entity.Metric) error
}

func (h handler) recordMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//
		start := time.Now()

		//next
		next.ServeHTTP(w, r)

		//
		duration := time.Since(start).Seconds()

		//metric
		//count request
		h.metric.RecordMetric(r.Context(), entity.Metric{
			Type:  entity.MetricTypeCounter,
			Name:  "http_requests_total",
			Value: 1,
			Labels: map[string]string{
				"method":   r.Method,
				"endpoint": r.URL.Path,
			},
		})

		//request processing time
		h.metric.RecordMetric(r.Context(), entity.Metric{
			Type:  entity.MetricTypeHistogram,
			Name:  "http_request_duration_seconds",
			Value: duration,
			Labels: map[string]string{
				"method":   r.Method,
				"endpoint": r.URL.Path,
			},
		})

	})
}
