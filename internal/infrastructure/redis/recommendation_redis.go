package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/probuborka/NutriAI/internal/entity"
	"github.com/redis/go-redis/v9"
)

type redisRecommendation struct {
	client *redis.Client
}

func NewRecommendation(client *redis.Client) *redisRecommendation {
	return &redisRecommendation{client: client}
}

// Finding recommendation in Redis
func (r *redisRecommendation) FindByIDNew(ctx context.Context, id string) (entity.UserRecommendationRequest, error) {
	recommendationCache := entity.UserRecommendationRequest{}

	data, err := r.client.Get(ctx, fmt.Sprintf("userID:%s", id)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return recommendationCache, nil
		}
		return recommendationCache, err
	}

	err = json.Unmarshal(data, &recommendationCache)
	if err != nil {
		return recommendationCache, err
	}

	return recommendationCache, nil
}

// Save recommendation in Redis
func (r *redisRecommendation) SaveNew(ctx context.Context, recommendation entity.UserRecommendationRequest) error {
	data, err := json.Marshal(recommendation)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, fmt.Sprintf("userID:%s", recommendation.UserID), string(data), 0).Err()
}
