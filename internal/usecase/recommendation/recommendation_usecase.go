package recommendation

import (
	"context"

	"github.com/probuborka/NutriAI/internal/entity"
)

type ai interface {
	Recommendation(userNFP entity.UserNutritionAndFitnessProfile) (string, error)
}

type cache interface {
	Save(ctx context.Context, recommendation entity.UserNutritionAndFitnessProfileCache) error
	FindByID(ctx context.Context, id string) (entity.UserNutritionAndFitnessProfileCache, error)
}

type service struct {
	ai    ai
	cache cache
}

func NewRecommendationUseCase(ai ai, cache cache) service {
	return service{
		ai:    ai,
		cache: cache,
	}
}

func (s service) GetRecommendation(ctx context.Context, userNFP entity.UserNutritionAndFitnessProfile) (string, error) {

	recommendationCache, err := s.cache.FindByID(ctx, userNFP.UserID)
	if err != nil {
		return "", err
	}

	if recommendationCache.Recommendations != "" &&
		recommendationCache.UserID == userNFP.UserID &&
		recommendationCache.Age == userNFP.Age &&
		recommendationCache.Gender == userNFP.Gender &&
		recommendationCache.Height == userNFP.Height &&
		recommendationCache.CurrentWeight == userNFP.CurrentWeight &&
		recommendationCache.GoalWeight == userNFP.GoalWeight &&
		recommendationCache.ActivityLevel == userNFP.ActivityLevel &&
		recommendationCache.DietaryPreferences == userNFP.DietaryPreferences &&
		recommendationCache.TrainingGoals == userNFP.TrainingGoals {
		return recommendationCache.Recommendations, err
	}

	str, err := s.ai.Recommendation(userNFP)
	if str == "" {
		return "", err
	}

	recommendation := entity.UserNutritionAndFitnessProfileCache{
		UserID:             userNFP.UserID,
		Age:                userNFP.Age,
		Gender:             userNFP.Gender,
		Height:             userNFP.Height,
		CurrentWeight:      userNFP.CurrentWeight,
		GoalWeight:         userNFP.GoalWeight,
		ActivityLevel:      userNFP.ActivityLevel,
		DietaryPreferences: userNFP.DietaryPreferences,
		TrainingGoals:      userNFP.TrainingGoals,
		Recommendations:    str,
	}

	err = s.cache.Save(ctx, recommendation)
	if err != nil {
		return "", err
	}

	return str, err
}
