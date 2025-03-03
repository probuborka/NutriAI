package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/probuborka/NutriAI/internal/entity"
	"github.com/redis/go-redis/v9"
)

// RedisTaskRepository — реализация TaskRepository для Redis
type redisRecommendation struct {
	client *redis.Client
}

// NewRedisTaskRepository — конструктор для RedisTaskRepository
func NewRecommendation(client *redis.Client) *redisRecommendation {
	return &redisRecommendation{client: client}
}

// Save — сохранение задачи в Redis
func (r *redisRecommendation) Save(ctx context.Context, recommendation entity.UserNutritionAndFitnessProfileCache) error {
	data, err := json.Marshal(recommendation)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, fmt.Sprintf("userID:%s", recommendation.UserID), string(data), 0).Err()
}

// FindByID — поиск задачи по ID в Redis
func (r *redisRecommendation) FindByID(ctx context.Context, id string) (entity.UserNutritionAndFitnessProfileCache, error) {
	recommendationCache := entity.UserNutritionAndFitnessProfileCache{}

	data, err := r.client.Get(ctx, fmt.Sprintf("userID:%s", id)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return recommendationCache, nil // Задача не найдена
		}
		return recommendationCache, err
	}

	err = json.Unmarshal(data, &recommendationCache)
	if err != nil {
		return recommendationCache, err
	}

	return recommendationCache, nil
}
