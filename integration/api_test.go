package integration

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/probuborka/NutriAI/internal/entity"
// 	"github.com/sirupsen/logrus"
// 	"github.com/stretchr/testify/assert"
// )

// // MockServiceRecommendation - мок для интерфейса serviceRecommendation
// type MockServiceRecommendation struct{}

// func (m *MockServiceRecommendation) GetRecommendation(ctx context.Context, userRecommendationRequest entity.UserRecommendationRequest) (string, error) {
// 	return "Eat more protein", nil
// }

// // MockMetric - мок для интерфейса metric
// type MockMetric struct{}

// func (m *MockMetric) RecordMetrics(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		next.ServeHTTP(w, r)
// 	})
// }

// func TestGetRecommendation_Integration(t *testing.T) {
// 	// Создаем моки для сервиса рекомендаций и метрик
// 	mockService := &MockServiceRecommendation{}
// 	mockMetric := &MockMetric{}

// 	// Создаем логгер
// 	logger := logrus.New()

// 	// Создаем обработчик
// 	handler := New(mockService, mockMetric, logger)

// 	// Инициализируем маршруты
// 	router := handler.Init()

// 	// Тестовые данные
// 	validRequest := entity.UserRecommendationRequest{
// 		UserID: "user123",
// 		UserData: entity.UserData{
// 			Profile: entity.Profile{
// 				Age:          30,
// 				Gender:       "female",
// 				WeightKg:     70,
// 				HeightCm:     165,
// 				FitnessLevel: "intermediate",
// 			},
// 			Goals: entity.Goals{
// 				PrimaryGoal:    "weight_loss",
// 				SecondaryGoal:  "muscle_toning",
// 				TargetWeightKg: 65,
// 				TimeframeWeeks: 12,
// 			},
// 			Preferences: entity.Preferences{
// 				DietType:           "balanced",
// 				Allergies:          []string{"nuts"},
// 				PreferredCuisines:  []string{"mediterranean"},
// 				WorkoutPreferences: []string{"yoga"},
// 			},
// 			Lifestyle: entity.Lifestyle{
// 				ActivityLevel:           "moderate",
// 				DailyCalorieIntake:      1800,
// 				WorkoutAvailabilityDays: 4,
// 				AverageSleepHours:       7,
// 			},
// 			MedicalRestrictions: entity.MedicalRestrictions{
// 				HasInjuries:       true,
// 				InjuryDetails:     []string{"lower_back_pain"},
// 				ChronicConditions: []string{"none"},
// 			},
// 		},
// 		RequestDetails: entity.RequestDetails{
// 			ServiceType:  "fitness_nutrition_recommendations",
// 			OutputFormat: "weekly_plan",
// 			Language:     "ru",
// 		},
// 	}

// 	t.Run("Success - valid request", func(t *testing.T) {
// 		// Преобразуем запрос в JSON
// 		requestBody, _ := json.Marshal(validRequest)

// 		// Создаем HTTP-запрос
// 		req := httptest.NewRequest(http.MethodGet, "/api/recommendation", bytes.NewBuffer(requestBody))
// 		req.Header.Set("Content-Type", "application/json")

// 		// Создаем ResponseWriter для записи ответа
// 		rec := httptest.NewRecorder()

// 		// Вызываем обработчик
// 		router.ServeHTTP(rec, req)

// 		// Проверяем статус код
// 		assert.Equal(t, http.StatusCreated, rec.Code)

// 		// Проверяем тело ответа
// 		var response entity.RecommendationResponse
// 		err := json.Unmarshal(rec.Body.Bytes(), &response)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Eat more protein", response.Recommendations)
// 	})

// 	// t.Run("Error - invalid JSON", func(t *testing.T) {
// 	// 	// Создаем HTTP-запрос с некорректным JSON
// 	// 	req := httptest.NewRequest(http.MethodGet, "/api/recommendation", bytes.NewBuffer([]byte("{invalid json}")))
// 	// 	req.Header.Set("Content-Type", "application/json")

// 	// 	// Создаем ResponseWriter для записи ответа
// 	// 	rec := httptest.NewRecorder()

// 	// 	// Вызываем обработчик
// 	// 	router.ServeHTTP(rec, req)

// 	// 	// Проверяем статус код
// 	// 	assert.Equal(t, http.StatusBadRequest, rec.Code)

// 	// 	// Проверяем тело ответа
// 	// 	var response entity.Error
// 	// 	err := json.Unmarshal(rec.Body.Bytes(), &response)
// 	// 	assert.NoError(t, err)
// 	// 	assert.Contains(t, response.Error, "unmarshal error")
// 	// })

// 	// t.Run("Error - validation failed", func(t *testing.T) {
// 	// 	// Создаем некорректный запрос (без UserID)
// 	// 	invalidRequest := validRequest
// 	// 	invalidRequest.UserID = ""

// 	// 	// Преобразуем запрос в JSON
// 	// 	requestBody, _ := json.Marshal(invalidRequest)

// 	// 	// Создаем HTTP-запрос
// 	// 	req := httptest.NewRequest(http.MethodGet, "/api/recommendation", bytes.NewBuffer(requestBody))
// 	// 	req.Header.Set("Content-Type", "application/json")

// 	// 	// Создаем ResponseWriter для записи ответа
// 	// 	rec := httptest.NewRecorder()

// 	// 	// Вызываем обработчик
// 	// 	router.ServeHTTP(rec, req)

// 	// 	// Проверяем статус код
// 	// 	assert.Equal(t, http.StatusBadRequest, rec.Code)

// 	// 	// Проверяем тело ответа
// 	// 	var response entity.Error
// 	// 	err := json.Unmarshal(rec.Body.Bytes(), &response)
// 	// 	assert.NoError(t, err)
// 	// 	assert.Contains(t, response.Error, "validation failed")
// 	// })
// }
