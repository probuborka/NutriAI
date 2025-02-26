package recommendation

import (
	"context"

	"github.com/probuborka/NutriAI/internal/entity"
)

type ai interface {
	Recommendation(userNFP entity.UserNutritionAndFitnessProfile) (string, error)
}

type cache interface {
	Save(ctx context.Context, id string, recommendation string) error
	FindByID(ctx context.Context, id string) (string, error)
}

type service struct {
	ai    ai
	cache cache
}

func New(ai ai, cache cache) service {
	return service{
		ai:    ai,
		cache: cache,
	}
}

func (s service) GetRecommendation(ctx context.Context, userNFP entity.UserNutritionAndFitnessProfile) (string, error) {

	str, err := s.cache.FindByID(ctx, userNFP.UserID)
	if str != "" {
		return str, err
	}

	str, err = s.ai.Recommendation(userNFP)
	if str != "" {
		s.cache.Save(ctx, userNFP.UserID, str)
	}
	return str, err
}
