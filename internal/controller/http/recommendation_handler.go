package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/probuborka/NutriAI/internal/entity"
	"github.com/sirupsen/logrus"
)

type serviceRecommendation interface {
	GetRecommendation(ctx context.Context, userRecommendationRequest entity.UserRecommendationRequest) (string, error)
}

// @Summary Получить рекомендации
// @Description Возвращает персонализированные рекомендации по питанию и тренировкам
// @Tags Рекомендации
// @Accept json
// @Produce json
// @Param request body entity.UserRecommendationRequest true "Данные пользователя"
// @Success 201 {object} entity.RecommendationResponse
// @Failure 400 {object} entity.Error
// @Router /api/recommendation [get]
func (h handler) getRecommendation(w http.ResponseWriter, r *http.Request) {
	//
	requestID, ok := r.Context().Value(requestIDKey).(string)
	if !ok {
		requestID = "unknown"
	}

	//
	var userRecommendationRequest entity.UserRecommendationRequest
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		h.response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest, requestID)
		h.log.WithFields(logrus.Fields{
			"requestID": requestID,
			"error":     err,
		}).Error("buf ReadFrom")
		return
	}

	err = json.Unmarshal(buf.Bytes(), &userRecommendationRequest)
	if err != nil {
		h.response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest, requestID)
		h.log.WithFields(logrus.Fields{
			"requestID": requestID,
			"error":     err,
		}).Error("unmarshal error")
		return
	}

	recommendations, err := h.recommendation.GetRecommendation(r.Context(), userRecommendationRequest)
	if err != nil {
		h.response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest, requestID)
		h.log.WithFields(logrus.Fields{
			"requestID": requestID,
			"error":     err,
		}).Error("usecase recommendations")
		return
	}

	//
	h.response(w, entity.RecommendationResponse{Recommendations: recommendations}, http.StatusCreated, requestID)
}
