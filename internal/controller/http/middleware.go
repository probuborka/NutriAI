package http

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

// logging
type bodyLogWriter struct {
	http.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (h handler) logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//
		requestId := uuid.New().String()

		//
		startTime := time.Now()

		// Логируем входящий запрос
		var requestBody string
		if req.Body != nil {
			bodyBytes, _ := io.ReadAll(req.Body)
			requestBody = string(bodyBytes)
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Восстанавливаем тело запроса
		}

		h.log.WithFields(logrus.Fields{
			"requestId": requestId,
			"method":    req.Method,
			"path":      req.URL.Path,
			"ip":        req.RemoteAddr,
			"body":      requestBody,
		}).Info("Incoming request")

		// Перехват ответа
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: w}
		w = blw

		// Передаем управление следующему обработчику
		next.ServeHTTP(w, req)

		// Логируем ответ
		duration := time.Since(startTime)
		h.log.WithFields(logrus.Fields{
			"requestId": requestId,
			"status":    blw.status,
			"duration":  duration,
			"response":  blw.body.String(),
		}).Info("Outgoing response")

	})
}
