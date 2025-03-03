package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// requestIDKey for storing in context
type contextKey string

const requestIDKey contextKey = "requestID"

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

func (w *bodyLogWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (h handler) logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//add RequestID to context
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		r = r.WithContext(ctx)

		//start
		start := time.Now()

		//log request
		var requestBody string
		if r.Body != nil {
			bodyBytes, _ := io.ReadAll(r.Body)
			requestBody = string(bodyBytes)
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) //restoring the request body
		}

		h.log.WithFields(logrus.Fields{
			"requestID": requestID,
			"method":    r.Method,
			"path":      r.URL.Path,
			"ip":        r.RemoteAddr,
			"body":      requestBody,
		}).Info("Incoming request")

		//
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: w,
		}
		w = blw

		//next
		next.ServeHTTP(w, r)

		//log response
		duration := time.Since(start).Seconds()

		//
		h.log.WithFields(logrus.Fields{
			"requestID": requestID,
			"status":    blw.status,
			"duration":  duration,
			"response":  blw.body.String(),
		}).Info("Outgoing response")

	})
}
