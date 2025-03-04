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

func (gch gigachatRecommendation) Recommendation(userRecommendation entity.UserRecommendationRequest) (string, error) {

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
