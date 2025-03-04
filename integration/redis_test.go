package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/probuborka/NutriAI/internal/infrastructure/redis"
	redisclient "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedis_Integration(t *testing.T) {

	t.Run("Redis Save and FindByID - valid", func(t *testing.T) {
		//client
		redisClient := redisclient.NewClient(
			&redisclient.Options{
				Addr:     fmt.Sprintf("%s:%s", RedisHost, RedisPort),
				Password: "",
				DB:       0,
			},
		)

		err := redisClient.Del(context.Background(), fmt.Sprintf("userID:%s", userRecommendationRequest.UserID)).Err()
		assert.NoError(t, err)
		//
		cacheRecommendation := redis.NewRecommendation(
			redisClient,
		)

		//Save
		err = cacheRecommendation.Save(context.Background(), userRecommendationRequest)
		assert.NoError(t, err)

		//FindByID
		recommendationRequest, err := cacheRecommendation.FindByID(context.Background(), userRecommendationRequest.UserID)
		assert.NoError(t, err)
		assert.NotEmpty(t, recommendationRequest.UserID)
	})

}
