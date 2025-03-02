package http

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h handler) response(w http.ResponseWriter, v any, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	resp, err := json.Marshal(&v)
	if err != nil {
		h.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("marshal error")
		return
	}
	w.Write(resp)
}
