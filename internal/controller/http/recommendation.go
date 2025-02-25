package http

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/probuborka/NutriAI/internal/entity"
	"github.com/probuborka/NutriAI/pkg/logger"
)

type serviceRecommendation interface {
	GetRecommNutriAL(userNFP entity.UserNutritionAndFitnessProfile) (string, error)
}

func (h handler) getRecommendation(w http.ResponseWriter, r *http.Request) {

	var userNFP entity.UserNutritionAndFitnessProfile
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		//
		//response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &userNFP)
	if err != nil {
		//
		//response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	str, err := h.recommendation.GetRecommNutriAL(userNFP)
	if err != nil {
		//
		//response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	_ = str

	//
	//response(w, entitytask.IdTask{ID: strconv.Itoa(id)}, http.StatusCreated)
}
