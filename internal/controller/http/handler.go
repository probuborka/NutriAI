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
	// authentication serviceAuthentication
}

func New(recommendation serviceRecommendation, metric metric, log *logrus.Logger) *handler {
	//cfg = cfgAuth
	return &handler{
		recommendation: recommendation,
		metric:         metric,
		log:            log,
		// authentication: authentication,
	}
}

func (h handler) Init() http.Handler {
	r := http.NewServeMux()

	// HTTP-обработчик для метрик
	r.Handle("/metrics", promhttp.Handler())

	//web
	//r.Handle("/", http.FileServer(http.Dir(entityconfig.WebDir)))

	//recommendation
	r.HandleFunc("GET /api/recommendation", h.getRecommendation)

	//create task
	// r.HandleFunc("POST /api/task", h.createTask)

	// //get tasks
	// r.HandleFunc("GET /api/tasks", h.getTasks)

	// //get task
	// r.HandleFunc("GET /api/task", h.getTask)

	// //change task
	// r.HandleFunc("PUT /api/task", h.changeTask)

	// //done task
	// r.HandleFunc("POST /api/task/done", h.doneTask)

	// //delete task
	// r.HandleFunc("DELETE /api/task", h.deleteTask)

	// //authentication
	// r.HandleFunc("POST /api/signin", h.password)

	//
	stack := []middleware{
		h.recordMetrics,
		h.logging,
	}

	hand := compileMiddleware(r, stack)

	return hand
	//return r
}
