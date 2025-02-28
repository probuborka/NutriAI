package metric

import (
	"context"

	"github.com/probuborka/NutriAI/internal/entity"
)

type metric interface {
	Save(ctx context.Context, metric entity.Metric) error
}

type metricUseCase struct {
	metric metric
}

func NewMetricUseCase(metric metric) *metricUseCase {
	return &metricUseCase{metric: metric}
}

func (uc *metricUseCase) RecordMetric(ctx context.Context, metric entity.Metric) error {
	return uc.metric.Save(ctx, metric)
}
