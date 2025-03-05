package http

import (
	"net/http"

	_ "github.com/probuborka/NutriAI/docs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
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

	//Swagger UI
	r.Handle("/swagger/", httpSwagger.WrapHandler)

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
