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

func TestGetRecommendation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAI := mocks.NewMockai(ctrl)
	mockCache := mocks.NewMockcache(ctrl)

	service := NewRecommendationUseCase(mockAI, mockCache)

	ctx := context.Background()

	validUserRecommendationRequest := entity.UserRecommendationRequest{
		UserID:   "123456789",
		UserName: "Евгений",
		UserData: entity.UserData{
			Profile: entity.Profile{
				Age:          39,
				Gender:       "male",
				WeightKg:     140,
				HeightCm:     186,
				FitnessLevel: "beginner",
			},
			Goals: entity.Goals{
				PrimaryGoal:    "weight_loss",
				SecondaryGoal:  "muscle_toning",
				TargetWeightKg: 90,
				TimeframeWeeks: 40,
			},
			Preferences: entity.Preferences{
				DietType:           "balanced",
				Allergies:          []string{"орехи", "моллюски"},
				PreferredCuisines:  []string{"средиземноморский", "азиатский"},
				WorkoutPreferences: []string{"йога", "силовая тренировка", "кардио"},
			},
			Lifestyle: entity.Lifestyle{
				ActivityLevel:           "moderate",
				DailyCalorieIntake:      1800,
				WorkoutAvailabilityDays: 4,
				AverageSleepHours:       7,
			},
			MedicalRestrictions: entity.MedicalRestrictions{
				HasInjuries:       true,
				InjuryDetails:     []string{"травма колена"},
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

		mockCache.EXPECT().FindByID(ctx, validUserRecommendationRequest.UserID).Return(recommendationCache, nil)

		result, err := service.GetRecommendation(ctx, validUserRecommendationRequest)
		assert.NoError(t, err)
		assert.Equal(t, "Eat more protein", result)
	})

	t.Run("Success - recommendation from AI", func(t *testing.T) {
		recommendationCache := entity.UserRecommendationRequest{
			UserID: validUserRecommendationRequest.UserID,
		}

		mockCache.EXPECT().FindByID(ctx, validUserRecommendationRequest.UserID).Return(recommendationCache, nil)
		mockAI.EXPECT().Recommendation(validUserRecommendationRequest).Return("Drink more water", nil)
		mockCache.EXPECT().Save(ctx, gomock.Any()).Return(nil)

		result, err := service.GetRecommendation(ctx, validUserRecommendationRequest)
		assert.NoError(t, err)
		assert.Equal(t, "Drink more water", result)
	})

	t.Run("Error - validation failed", func(t *testing.T) {
		invalidUserRecommendationRequest := entity.UserRecommendationRequest{
			UserID: "", // Invalid UserID
		}

		result, err := service.GetRecommendation(ctx, invalidUserRecommendationRequest)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("Error - cache FindByID failed", func(t *testing.T) {
		mockCache.EXPECT().FindByID(ctx, validUserRecommendationRequest.UserID).Return(entity.UserRecommendationRequest{}, errors.New("cache error"))

		result, err := service.GetRecommendation(ctx, validUserRecommendationRequest)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("Error - AI recommendation failed", func(t *testing.T) {
		recommendationCache := entity.UserRecommendationRequest{
			UserID: validUserRecommendationRequest.UserID,
		}

		mockCache.EXPECT().FindByID(ctx, validUserRecommendationRequest.UserID).Return(recommendationCache, nil)
		mockAI.EXPECT().Recommendation(validUserRecommendationRequest).Return("", errors.New("AI error"))

		result, err := service.GetRecommendation(ctx, validUserRecommendationRequest)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("Error - cache Save failed", func(t *testing.T) {
		recommendationCache := entity.UserRecommendationRequest{
			UserID: validUserRecommendationRequest.UserID,
		}

		mockCache.EXPECT().FindByID(ctx, validUserRecommendationRequest.UserID).Return(recommendationCache, nil)
		mockAI.EXPECT().Recommendation(validUserRecommendationRequest).Return("Drink more water", nil)
		mockCache.EXPECT().Save(ctx, gomock.Any()).Return(errors.New("save error"))

		result, err := service.GetRecommendation(ctx, validUserRecommendationRequest)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})
}
