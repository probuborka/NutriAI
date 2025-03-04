package gigachat

import (
	"encoding/json"
	"fmt"

	"github.com/probuborka/NutriAI/internal/entity"
	"github.com/probuborka/NutriAI/pkg/gigachat"
)

type gigachatRecommendation struct {
	client *gigachat.Client
}

func NewRecommendation(client *gigachat.Client) *gigachatRecommendation {
	return &gigachatRecommendation{
		client: client,
	}
}

func (gch gigachatRecommendation) Recommendation(userNFP entity.UserNutritionAndFitnessProfile) (string, error) {

	contentSystem := `Отвечай как нутрициолог`

	contentUser := fmt.Sprintf(`
		Дай персонализированные рекомендации на основе данных от пользователя: 
		Пол: %s, 
		Возраст: %v лет, 
		Рост: %v см, 
		Текущий вес: %v, 
		Желаемый вес: %v, 
		Уровень физической активности пользователя: %s, 
		Предпочтения в питании: %s
		Цель: %s`,
		userNFP.Gender,             // Пол пользователя
		userNFP.Age,                // Возраст пользователя
		userNFP.Height,             // Рост пользователя
		userNFP.CurrentWeight,      // Текущий вес пользователя
		userNFP.GoalWeight,         // Желаемый вес пользователя
		userNFP.ActivityLevel,      // Уровень физической активности пользователя
		userNFP.DietaryPreferences, // Предпочтения в питании
		userNFP.TrainingGoals,      // Цели тренировок
	)

	// message
	message := gigachat.RequestBody{
		Model:           "GigaChat",
		Stream:          false,
		Update_interval: 0,
		Messages: []gigachat.Messages{
			{
				Role:    "system",
				Content: contentSystem,
			},
			{
				Role:    "user",
				Content: contentUser,
			},
		},
	}

	result, err := gch.client.GenerateText(message)
	if err != nil {
		return "", err
	}

	str := ""
	for _, v := range result.Choices {
		str = v.Message.Content
		break
	}

	return str, err

}

// new
func (gch gigachatRecommendation) RecommendationNew(userRecommendation entity.UserRecommendationRequest) (string, error) {

	data, err := json.Marshal(&userRecommendation)
	if err != nil {
		return "", err
	}

	contentSystem := `Отвечай как нутрициолог`

	contentUser := fmt.Sprintf(`
		Дай персонализированные рекомендации на основе данных от пользователя: 
		%s`,
		string(data),
	)

	// message
	message := gigachat.RequestBody{
		Model:           "GigaChat",
		Stream:          false,
		Update_interval: 0,
		Messages: []gigachat.Messages{
			{
				Role:    "system",
				Content: contentSystem,
			},
			{
				Role:    "user",
				Content: contentUser,
			},
		},
	}

	result, err := gch.client.GenerateText(message)
	if err != nil {
		return "", err
	}

	str := ""
	for _, v := range result.Choices {
		str = v.Message.Content
		break
	}

	return str, err

}
