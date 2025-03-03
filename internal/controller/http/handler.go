package http

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

type handler struct {
	recommendation serviceRecommendation
	metric         metric
	log            *logrus.Logger
}

func New(recommendation serviceRecommendation, metric metric, log *logrus.Logger) *handler {
	return &handler{
		recommendation: recommendation,
		metric:         metric,
		log:            log,
	}
}

func (h handler) Init() http.Handler {
	r := http.NewServeMux()

	//metrics
	r.Handle("/metrics", promhttp.Handler())

	//recommendation
	r.HandleFunc("GET /api/recommendation", h.getRecommendation)

	//middleware
	stack := []middleware{
		h.recordMetrics,
		h.logging,
	}

	return compileMiddleware(r, stack)
}
