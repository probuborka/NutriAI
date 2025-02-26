package redis

import (
	"context"
	"errors"
	"fmt"

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
func (r *redisRecommendation) Save(ctx context.Context, id string, recommendation string) error {
	// data, err := json.Marshal(task)
	// if err != nil {
	// 	return err
	// }
	return r.client.Set(ctx, fmt.Sprintf("userID:%s", id), recommendation, 0).Err()
}

// FindByID — поиск задачи по ID в Redis
func (r *redisRecommendation) FindByID(ctx context.Context, id string) (string, error) {
	data, err := r.client.Get(ctx, fmt.Sprintf("userID:%s", id)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil // Задача не найдена
		}
		return "", err
	}

	// var task domain.Task
	// if err := json.Unmarshal(data, &task); err != nil {
	// 	return nil, err
	// }
	return string(data), nil
}
