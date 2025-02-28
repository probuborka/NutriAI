package entity

type MetricType string

const (
	MetricTypeCounter   MetricType = "counter"
	MetricTypeHistogram MetricType = "histogram"
)

type Metric struct {
	Type   MetricType
	Name   string
	Value  float64
	Labels map[string]string
}
