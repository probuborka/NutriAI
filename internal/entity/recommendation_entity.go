package entity

import "errors"

// Описание полей:
// Age: Возраст пользователя.
// Gender: Пол пользователя.
// Height: Рост пользователя.
// CurrentWeight: Текущий вес пользователя.
// GoalWeight: Желаемый вес пользователя.
// ActivityLevel: Уровень физической активности пользователя.
// DietaryPreferences: Предпочтения в питании.
// TrainingGoals: Цели тренировок.
type UserNutritionAndFitnessProfile struct {
	UserID             string  `json:"user_id"`
	Age                int     `json:"age"`
	Gender             string  `json:"gender"`
	Height             float32 `json:"height"`
	CurrentWeight      float32 `json:"current_weight"`
	GoalWeight         float32 `json:"goal_weight"`
	ActivityLevel      string  `json:"activity_level"`
	DietaryPreferences string  `json:"dietary_preferences"`
	TrainingGoals      string  `json:"training_goals"`
}

func (u UserNutritionAndFitnessProfile) Validate() error {
	//UserID
	if u.UserID == "" {
		return errors.New("user_id is required")
	}
	if len(u.UserID) < 3 || len(u.UserID) > 50 {
		return errors.New("user_id must be between 3 and 50 characters")
	}

	//Age
	if u.Age < 18 || u.Age > 100 {
		return errors.New("age must be between 18 and 100")
	}

	//Gender
	if u.Gender == "" {
		return errors.New("gender is required")
	}
	if u.Gender != "male" && u.Gender != "female" {
		return errors.New("gender is error. (male or female)")
	}

	//Height
	if u.Height == 0 {
		return errors.New("height is required")
	}
	if u.Height < 50 || u.Height > 250 {
		return errors.New("height must be between 50 and 250 cm")
	}

	//CurrentWeight
	if u.CurrentWeight == 0 {
		return errors.New("currentWeight is required")
	}

	//GoalWeight
	if u.GoalWeight == 0 {
		return errors.New("goalWeight is required")
	}

	//ActivityLevel
	if u.ActivityLevel == "" {
		return errors.New("activityLevel is required")
	}

	//DietaryPreferences
	if u.DietaryPreferences == "" {
		return errors.New("dietaryPreferences is required")
	}

	//TrainingGoals
	if u.TrainingGoals == "" {
		return errors.New("trainingGoals is required")
	}
	// if !strings.Contains(u.Email, "@") {
	//     return errors.New("invalid email")
	// }
	return nil
}

type RecommendationResponse struct {
	Recommendations string `json:"recommendations"`
}

type UserNutritionAndFitnessProfileCache struct {
	UserID             string  `json:"user_id"`
	Age                int     `json:"age"`
	Gender             string  `json:"gender"`
	Height             float32 `json:"height"`
	CurrentWeight      float32 `json:"current_weight"`
	GoalWeight         float32 `json:"goal_weight"`
	ActivityLevel      string  `json:"activity_level"`
	DietaryPreferences string  `json:"dietary_preferences"`
	TrainingGoals      string  `json:"training_goals"`
	Recommendations    string  `json:"recommendations"`
}
