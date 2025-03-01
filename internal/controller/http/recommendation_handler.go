package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/probuborka/NutriAI/internal/entity"
	"github.com/probuborka/NutriAI/pkg/logger"
)

type serviceRecommendation interface {
	GetRecommendation(ctx context.Context, userNFP entity.UserNutritionAndFitnessProfile) (string, error)
}

func (h handler) getRecommendation(w http.ResponseWriter, r *http.Request) {

	var userNFP entity.UserNutritionAndFitnessProfile
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest)
		logger.Error(err)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &userNFP)
	if err != nil {
		response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest)
		logger.Error(err)
		return
	}

	recommendations, err := h.recommendation.GetRecommendation(r.Context(), userNFP)
	if err != nil {
		response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest)
		logger.Error(err)
		return
	}

	//
	response(w, entity.RecommendationResponse{Recommendations: recommendations}, http.StatusCreated)
}
