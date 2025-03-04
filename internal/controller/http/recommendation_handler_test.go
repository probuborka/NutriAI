package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/probuborka/NutriAI/internal/entity"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockServiceRecommendation - мок для интерфейса serviceRecommendation
type MockServiceRecommendation struct {
	mock.Mock
}

func (m *MockServiceRecommendation) GetRecommendation(ctx context.Context, userRecommendationRequest entity.UserRecommendationRequest) (string, error) {
	args := m.Called(ctx, userRecommendationRequest)
	return args.String(0), args.Error(1)
}

func TestGetRecommendation(t *testing.T) {
	// Создаем мок для serviceRecommendation
	mockService := new(MockServiceRecommendation)

	// Создаем логгер
	logger := logrus.New()

	// Создаем обработчик
	handler := handler{
		recommendation: mockService,
		log:            logger,
	}

	// Тестовые данные
	validRequest := entity.UserRecommendationRequest{
		UserID:   "user123",
		UserName: "jenya",
		UserData: entity.UserData{
			Profile: entity.Profile{
				Age:          30,
				Gender:       "female",
				WeightKg:     70,
				HeightCm:     165,
				FitnessLevel: "intermediate",
			},
			Goals: entity.Goals{
				PrimaryGoal:    "weight_loss",
				SecondaryGoal:  "muscle_toning",
				TargetWeightKg: 65,
				TimeframeWeeks: 12,
			},
			Preferences: entity.Preferences{
				DietType:           "balanced",
				Allergies:          []string{"nuts"},
				PreferredCuisines:  []string{"mediterranean"},
				WorkoutPreferences: []string{"yoga"},
			},
			Lifestyle: entity.Lifestyle{
				ActivityLevel:           "moderate",
				DailyCalorieIntake:      1800,
				WorkoutAvailabilityDays: 4,
				AverageSleepHours:       7,
			},
			MedicalRestrictions: entity.MedicalRestrictions{
				HasInjuries:       true,
				InjuryDetails:     []string{"lower_back_pain"},
				ChronicConditions: []string{"none"},
			},
		},
		RequestDetails: entity.RequestDetails{
			ServiceType:  "fitness_nutrition_recommendations",
			OutputFormat: "weekly_plan",
			Language:     "ru",
		},
	}

	t.Run("Success - valid request", func(t *testing.T) {
		// Ожидаем вызов метода GetRecommendationNew с корректными данными
		mockService.On("GetRecommendationNew", mock.Anything, validRequest).Return("Eat more protein", nil)

		// Преобразуем запрос в JSON
		requestBody, _ := json.Marshal(validRequest)

		// Создаем HTTP-запрос
		req := httptest.NewRequest(http.MethodGet, "/api/recommendationnew", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Создаем ResponseWriter для записи ответа
		rec := httptest.NewRecorder()

		// Вызываем обработчик
		handler.getRecommendationNew(rec, req)

		// Проверяем статус код
		assert.Equal(t, http.StatusCreated, rec.Code)

		// Проверяем тело ответа
		var response entity.RecommendationResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Eat more protein", response.Recommendations)

		// Проверяем, что метод мока был вызван
		mockService.AssertCalled(t, "GetRecommendationNew", mock.Anything, validRequest)
	})

	t.Run("Error - invalid JSON", func(t *testing.T) {
		// Создаем HTTP-запрос с некорректным JSON
		req := httptest.NewRequest(http.MethodGet, "/api/recommendationnew", bytes.NewBuffer([]byte("{invalid json}")))
		req.Header.Set("Content-Type", "application/json")

		// Создаем ResponseWriter для записи ответа
		rec := httptest.NewRecorder()

		// Вызываем обработчик
		handler.getRecommendationNew(rec, req)

		// Проверяем статус код
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		// Проверяем тело ответа
		var response entity.Error
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		//assert.Contains(t, response.Error, "unmarshal error")
	})

	t.Run("Error - validation failed", func(t *testing.T) {
		// Создаем некорректный запрос (без UserID)
		invalidRequest := validRequest
		invalidRequest.UserID = ""

		// Преобразуем запрос в JSON
		requestBody, _ := json.Marshal(invalidRequest)

		// Создаем HTTP-запрос
		req := httptest.NewRequest(http.MethodGet, "/api/recommendationnew", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Создаем ResponseWriter для записи ответа
		rec := httptest.NewRecorder()

		// Вызываем обработчик
		handler.getRecommendationNew(rec, req)

		// Проверяем статус код
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		// Проверяем тело ответа
		var response entity.Error
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response.Error, "validation failed")
	})

	t.Run("Error - service error", func(t *testing.T) {
		// Ожидаем вызов метода GetRecommendationNew с ошибкой
		mockService.On("GetRecommendationNew", mock.Anything, validRequest).Return("", errors.New("service error"))

		// Преобразуем запрос в JSON
		requestBody, _ := json.Marshal(validRequest)

		// Создаем HTTP-запрос
		req := httptest.NewRequest(http.MethodGet, "/api/recommendationnew", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Создаем ResponseWriter для записи ответа
		rec := httptest.NewRecorder()

		// Вызываем обработчик
		handler.getRecommendationNew(rec, req)

		// Проверяем статус код
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		// Проверяем тело ответа
		var response entity.Error
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "service error", response.Error)
	})
}
