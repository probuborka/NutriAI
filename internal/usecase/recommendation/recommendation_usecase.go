package recommendation

import (
	"context"

	"github.com/probuborka/NutriAI/internal/entity"
)

type ai interface {
	Recommendation(userNFP entity.UserNutritionAndFitnessProfile) (string, error)
	RecommendationNew(userRecommendation entity.UserRecommendationRequest) (string, error)
}

type cache interface {
	Save(ctx context.Context, recommendation entity.UserNutritionAndFitnessProfileCache) error
	FindByID(ctx context.Context, id string) (entity.UserNutritionAndFitnessProfileCache, error)
	//new
	SaveNew(ctx context.Context, recommendation entity.UserRecommendationRequest) error
	FindByIDNew(ctx context.Context, id string) (entity.UserRecommendationRequest, error)
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

	//validate
	err := userNFP.Validate()
	if err != nil {
		return "", err
	}

	//search for recommendations in cache
	recommendationCache, err := s.cache.FindByID(ctx, userNFP.UserID)
	if err != nil {
		return "", err
	}

	//check recommendation from cache and userNFP
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
		return recommendationCache.Recommendations, nil
	}

	//get recommendations from AI
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

	//save recommendations in cache
	err = s.cache.Save(ctx, recommendation)
	if err != nil {
		return "", err
	}

	return str, nil
}

// new ----------------------------------------------------------------------------------------------------
func (s service) GetRecommendationNew(ctx context.Context, userRecommendationRequest entity.UserRecommendationRequest) (string, error) {

	//validate
	err := userRecommendationRequest.Validate()
	if err != nil {
		return "", err
	}

	//search for recommendations in cache
	recommendationCache, err := s.cache.FindByIDNew(ctx, userRecommendationRequest.UserID)
	if err != nil {
		return "", err
	}

	//check recommendation from cache and user recommendation from request
	if userRecommendationRequest.UserID == recommendationCache.UserID &&
		userRecommendationRequest.UserData.Profile == recommendationCache.UserData.Profile &&
		userRecommendationRequest.UserData.Goals == recommendationCache.UserData.Goals {
		return recommendationCache.Recommendations, nil
	}

	//get recommendations from AI
	recommendations, err := s.ai.RecommendationNew(userRecommendationRequest)
	if recommendations == "" || err != nil {
		return "", err
	}

	userRecommendationRequest.Recommendations = recommendations

	//save recommendations in cache
	err = s.cache.SaveNew(ctx, userRecommendationRequest)
	if err != nil {
		return "", err
	}

	return recommendations, nil
}
