package recommendation

import (
	"context"

	"github.com/probuborka/NutriAI/internal/entity"
)

type ai interface {
	RecommendationNew(userRecommendation entity.UserRecommendationRequest) (string, error)
}

type cache interface {
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
		userRecommendationRequest.UserName == recommendationCache.UserName &&
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
