package recommendation

import (
	"github.com/probuborka/NutriAI/internal/entity"
	"github.com/probuborka/NutriAI/pkg/gigachat"
)

type gigaChat interface {
	GenerateText(body gigachat.RequestBody) (error, gigachat.ChatCompletionResult)
}

type service struct {
	gigachat gigaChat
}

func New(gigachat gigaChat) service {
	return service{
		gigachat: gigachat,
	}
}

func (s service) GetRecommNutriAL(userNFP entity.UserNutritionAndFitnessProfile) (string, error) {

	// message
	message := gigachat.RequestBody{
		Model:           "GigaChat",
		Stream:          false,
		Update_interval: 0,
		Messages: []gigachat.Messages{
			{
				Role:    "system", // контекст
				Content: "Отвечай как нутрициолог",
			},
			{
				Role:    "user", // контекст
				Content: "Напиши 5 вариантов названий для космической станции",
			},
		},
	}

	s.gigachat.GenerateText(message)

	return "", nil
}
