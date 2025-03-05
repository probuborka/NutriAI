package integration

import "github.com/probuborka/NutriAI/internal/entity"

const (
	//GigaChat
	valid_api_key   = "ZDMxOTdmNjUtMmY3MS00MTdjLThkY2YtODljY2RiZGI1ZDZkOjVlMmM3OWYxLTUwNDQtNDRkNi05NTY1LTA3NzBlNTkyMWNmMQ=="
	invalid_api_key = "ZDMxOTdmNjUtMmY3MS00MTdjLThkY2YtODljY2RiZGI1ZDZkOmEwN2Q5YjhkLWVlNDAtNDUzZS04MTk1LTYzMDQxODU0NjYwMA=="

	//Redis
	RedisHost = "localhost"
	RedisPort = "6379"
)

var userRecommendationRequest = entity.UserRecommendationRequest{
	UserID:   "123456789",
	UserName: "Евгений",
	UserData: entity.UserData{
		Profile: entity.Profile{
			Age:          39,
			Gender:       "male",
			WeightKg:     140,
			HeightCm:     186,
			FitnessLevel: "beginner",
		},
		Goals: entity.Goals{
			PrimaryGoal:    "weight_loss",
			SecondaryGoal:  "muscle_toning",
			TargetWeightKg: 90,
			TimeframeWeeks: 40,
		},
		Preferences: entity.Preferences{
			DietType:           "balanced",
			Allergies:          []string{"орехи", "моллюски"},
			PreferredCuisines:  []string{"средиземноморский", "азиатский"},
			WorkoutPreferences: []string{"йога", "силовая тренировка", "кардио"},
		},
		Lifestyle: entity.Lifestyle{
			ActivityLevel:           "moderate",
			DailyCalorieIntake:      1800,
			WorkoutAvailabilityDays: 4,
			AverageSleepHours:       7,
		},
		MedicalRestrictions: entity.MedicalRestrictions{
			HasInjuries:       true,
			InjuryDetails:     []string{"травма колена"},
			ChronicConditions: []string{"none"},
		},
	},
	RequestDetails: entity.RequestDetails{
		ServiceType:  "fitness_nutrition_recommendations",
		OutputFormat: "weekly_plan",
		Language:     "ru",
	},
}
