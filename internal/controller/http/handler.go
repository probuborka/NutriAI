package http

import (
	"net/http"
)

type handler struct {
	recommendation serviceRecommendation
	// authentication serviceAuthentication
}

func New(recommendation serviceRecommendation) *handler {
	//cfg = cfgAuth
	return &handler{
		recommendation: recommendation,
		// authentication: authentication,
	}
}

func (h handler) Init() http.Handler {
	r := http.NewServeMux()

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
	// stack := []middleware{
	// 	logging,
	// 	authentication,
	// }

	// hand := compileMiddleware(r, stack)

	// return hand
	return r
}
