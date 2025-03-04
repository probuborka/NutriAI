package integration

import (
	"testing"

	"github.com/probuborka/NutriAI/internal/infrastructure/gigachat"
	gigachatclient "github.com/probuborka/NutriAI/pkg/gigachat"
	"github.com/stretchr/testify/assert"
)

func TestGigaChat_Integration(t *testing.T) {

	t.Run("GigaChat GetRecommendation - valid", func(t *testing.T) {
		//client
		gigaChatClient := gigachatclient.New(
			valid_api_key,
		)

		//gigachat
		gigaChatRecommendation := gigachat.NewRecommendation(
			gigaChatClient,
		)

		str, err := gigaChatRecommendation.Recommendation(userRecommendationRequest)
		assert.NoError(t, err)
		assert.NotEmpty(t, str)
	})

	t.Run("GigaChat GetRecommendation - authorization error", func(t *testing.T) {
		//client
		gigaChatClient := gigachatclient.New(
			invalid_api_key,
		)

		//gigachat
		gigaChatRecommendation := gigachat.NewRecommendation(
			gigaChatClient,
		)

		_, err := gigaChatRecommendation.Recommendation(userRecommendationRequest)
		assert.EqualError(t, err, gigachatclient.ErrorAuthorizationError.Error())
	})

	t.Run("GigaChat GetRecommendation - invalid api key", func(t *testing.T) {
		//client
		gigaChatClient := gigachatclient.New(
			"invalid_api_key",
		)

		//gigachat
		gigaChatRecommendation := gigachat.NewRecommendation(
			gigaChatClient,
		)

		_, err := gigaChatRecommendation.Recommendation(userRecommendationRequest)
		assert.EqualError(t, err, gigachatclient.ErrorBadRequest.Error())
	})
}
