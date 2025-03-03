package entity

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
