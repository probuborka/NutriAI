package integration

import (
	"testing"

	"github.com/probuborka/NutriAI/internal/entity"
	"github.com/probuborka/NutriAI/internal/infrastructure/gigachat"
	gigachatclient "github.com/probuborka/NutriAI/pkg/gigachat"
	"github.com/stretchr/testify/assert"
)

const (
	valid_api_key   = "ZDMxOTdmNjUtMmY3MS00MTdjLThkY2YtODljY2RiZGI1ZDZkOjVlMmM3OWYxLTUwNDQtNDRkNi05NTY1LTA3NzBlNTkyMWNmMQ=="
	invalid_api_key = "ZDMxOTdmNjUtMmY3MS00MTdjLThkY2YtODljY2RiZGI1ZDZkOmEwN2Q5YjhkLWVlNDAtNDUzZS04MTk1LTYzMDQxODU0NjYwMA=="
)

var userRecommendationRequest = entity.UserRecommendationRequest{
	UserID:   "12345",
	UserName: "jenya",
	UserData: entity.UserData{
		Profile: entity.Profile{
			Age:          30,
			Gender:       "female",
			WeightKg:     70,
			HeightCm:     165,
			FitnessLevel: "intermediate",
		},
		Goals: entity.Goals{
			PrimaryGoal:    "weight_loss",
			SecondaryGoal:  "muscle_toning",
			TargetWeightKg: 65,
			TimeframeWeeks: 12,
		},
		Preferences: entity.Preferences{
			DietType:           "balanced",
			Allergies:          []string{"nuts"},
			PreferredCuisines:  []string{"mediterranean"},
			WorkoutPreferences: []string{"yoga"},
		},
		Lifestyle: entity.Lifestyle{
			ActivityLevel:           "moderate",
			DailyCalorieIntake:      1800,
			WorkoutAvailabilityDays: 4,
			AverageSleepHours:       7,
		},
		MedicalRestrictions: entity.MedicalRestrictions{
			HasInjuries:       true,
			InjuryDetails:     []string{"lower_back_pain"},
			ChronicConditions: []string{"none"},
		},
	},
	RequestDetails: entity.RequestDetails{
		ServiceType:  "fitness_nutrition_recommendations",
		OutputFormat: "weekly_plan",
		Language:     "ru",
	},
}

func TestGetRecommendation_Integration(t *testing.T) {

	t.Run("GigaChat recommendation - valid", func(t *testing.T) {
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

	t.Run("GigaChat recommendation - authorization error", func(t *testing.T) {
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

	t.Run("GigaChat recommendation - invalid api key", func(t *testing.T) {
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
