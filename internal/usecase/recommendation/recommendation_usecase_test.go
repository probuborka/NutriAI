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

	validUserNFP := entity.UserNutritionAndFitnessProfile{
		UserID:             "user123",
		Age:                25,
		Gender:             "male",
		Height:             180,
		CurrentWeight:      70,
		GoalWeight:         65,
		ActivityLevel:      "active",
		DietaryPreferences: "vegan",
		TrainingGoals:      "strength",
	}

	t.Run("Success - recommendation from cache", func(t *testing.T) {
		recommendationCache := entity.UserNutritionAndFitnessProfileCache{
			UserID:             validUserNFP.UserID,
			Age:                validUserNFP.Age,
			Gender:             validUserNFP.Gender,
			Height:             validUserNFP.Height,
			CurrentWeight:      validUserNFP.CurrentWeight,
			GoalWeight:         validUserNFP.GoalWeight,
			ActivityLevel:      validUserNFP.ActivityLevel,
			DietaryPreferences: validUserNFP.DietaryPreferences,
			TrainingGoals:      validUserNFP.TrainingGoals,
			Recommendations:    "Eat more protein",
		}

		mockCache.EXPECT().FindByID(ctx, validUserNFP.UserID).Return(recommendationCache, nil)

		result, err := service.GetRecommendation(ctx, validUserNFP)
		assert.NoError(t, err)
		assert.Equal(t, "Eat more protein", result)
	})

	t.Run("Success - recommendation from AI", func(t *testing.T) {
		recommendationCache := entity.UserNutritionAndFitnessProfileCache{
			UserID: validUserNFP.UserID,
		}

		mockCache.EXPECT().FindByID(ctx, validUserNFP.UserID).Return(recommendationCache, nil)
		mockAI.EXPECT().Recommendation(validUserNFP).Return("Drink more water", nil)
		mockCache.EXPECT().Save(ctx, gomock.Any()).Return(nil)

		result, err := service.GetRecommendation(ctx, validUserNFP)
		assert.NoError(t, err)
		assert.Equal(t, "Drink more water", result)
	})

	t.Run("Error - validation failed", func(t *testing.T) {
		invalidUserNFP := entity.UserNutritionAndFitnessProfile{
			UserID: "", // Invalid UserID
		}

		result, err := service.GetRecommendation(ctx, invalidUserNFP)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("Error - cache FindByID failed", func(t *testing.T) {
		mockCache.EXPECT().FindByID(ctx, validUserNFP.UserID).Return(entity.UserNutritionAndFitnessProfileCache{}, errors.New("cache error"))

		result, err := service.GetRecommendation(ctx, validUserNFP)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("Error - AI recommendation failed", func(t *testing.T) {
		recommendationCache := entity.UserNutritionAndFitnessProfileCache{
			UserID: validUserNFP.UserID,
		}

		mockCache.EXPECT().FindByID(ctx, validUserNFP.UserID).Return(recommendationCache, nil)
		mockAI.EXPECT().Recommendation(validUserNFP).Return("", errors.New("AI error"))

		result, err := service.GetRecommendation(ctx, validUserNFP)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("Error - cache Save failed", func(t *testing.T) {
		recommendationCache := entity.UserNutritionAndFitnessProfileCache{
			UserID: validUserNFP.UserID,
		}

		mockCache.EXPECT().FindByID(ctx, validUserNFP.UserID).Return(recommendationCache, nil)
		mockAI.EXPECT().Recommendation(validUserNFP).Return("Drink more water", nil)
		mockCache.EXPECT().Save(ctx, gomock.Any()).Return(errors.New("save error"))

		result, err := service.GetRecommendation(ctx, validUserNFP)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})
}
