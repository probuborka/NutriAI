package http

import (
	"net/http"
)

type middleware func(http.Handler) http.Handler

func compileMiddleware(h http.Handler, m []middleware) http.Handler {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}
