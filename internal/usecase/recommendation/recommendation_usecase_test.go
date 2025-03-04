package recommendation

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/probuborka/NutriAI/internal/entity"
	"github.com/probuborka/NutriAI/internal/usecase/recommendation/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetRecommendationNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAI := mocks.NewMockai(ctrl)
	mockCache := mocks.NewMockcache(ctrl)

	service := NewRecommendationUseCase(mockAI, mockCache)

	ctx := context.Background()

	validUserRecommendationRequest := entity.UserRecommendationRequest{
		UserID: "user123",
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

	t.Run("Success - recommendation from cache", func(t *testing.T) {
		recommendationCache := entity.UserRecommendationRequest{
			UserID:   validUserRecommendationRequest.UserID,
			UserName: validUserRecommendationRequest.UserName,
			UserData: entity.UserData{
				Profile: validUserRecommendationRequest.UserData.Profile,
				Goals:   validUserRecommendationRequest.UserData.Goals,
			},
			Recommendations: "Eat more protein",
		}

		mockCache.EXPECT().FindByIDNew(ctx, validUserRecommendationRequest.UserID).Return(recommendationCache, nil)

		result, err := service.GetRecommendationNew(ctx, validUserRecommendationRequest)
		assert.NoError(t, err)
		assert.Equal(t, "Eat more protein", result)
	})

	t.Run("Success - recommendation from AI", func(t *testing.T) {
		recommendationCache := entity.UserRecommendationRequest{
			UserID: validUserRecommendationRequest.UserID,
		}

		mockCache.EXPECT().FindByIDNew(ctx, validUserRecommendationRequest.UserID).Return(recommendationCache, nil)
		mockAI.EXPECT().RecommendationNew(validUserRecommendationRequest).Return("Drink more water", nil)
		mockCache.EXPECT().SaveNew(ctx, gomock.Any()).Return(nil)

		result, err := service.GetRecommendationNew(ctx, validUserRecommendationRequest)
		assert.NoError(t, err)
		assert.Equal(t, "Drink more water", result)
	})

	t.Run("Error - validation failed", func(t *testing.T) {
		invalidUserRecommendationRequest := entity.UserRecommendationRequest{
			UserID: "", // Invalid UserID
		}

		result, err := service.GetRecommendationNew(ctx, invalidUserRecommendationRequest)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("Error - cache FindByIDNew failed", func(t *testing.T) {
		mockCache.EXPECT().FindByIDNew(ctx, validUserRecommendationRequest.UserID).Return(entity.UserRecommendationRequest{}, errors.New("cache error"))

		result, err := service.GetRecommendationNew(ctx, validUserRecommendationRequest)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("Error - AI recommendation failed", func(t *testing.T) {
		recommendationCache := entity.UserRecommendationRequest{
			UserID: validUserRecommendationRequest.UserID,
		}

		mockCache.EXPECT().FindByIDNew(ctx, validUserRecommendationRequest.UserID).Return(recommendationCache, nil)
		mockAI.EXPECT().RecommendationNew(validUserRecommendationRequest).Return("", errors.New("AI error"))

		result, err := service.GetRecommendationNew(ctx, validUserRecommendationRequest)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("Error - cache SaveNew failed", func(t *testing.T) {
		recommendationCache := entity.UserRecommendationRequest{
			UserID: validUserRecommendationRequest.UserID,
		}

		mockCache.EXPECT().FindByIDNew(ctx, validUserRecommendationRequest.UserID).Return(recommendationCache, nil)
		mockAI.EXPECT().RecommendationNew(validUserRecommendationRequest).Return("Drink more water", nil)
		mockCache.EXPECT().SaveNew(ctx, gomock.Any()).Return(errors.New("save error"))

		result, err := service.GetRecommendationNew(ctx, validUserRecommendationRequest)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})
}
