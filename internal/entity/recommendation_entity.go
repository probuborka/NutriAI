package entity

import (
	"errors"
	"fmt"
	"slices"

	"github.com/go-playground/validator"
)

// old
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

// ActivityLevels
// Sedentary
// Описание: Минимальная физическая активность или ее отсутствие (например, офисная работа без тренировок).
// Lightly Active
// Описание: Легкая физическая активность 1–3 раза в неделю (например, прогулки или легкие тренировки).
// Moderately Active
// Описание: Умеренная физическая активность 3–5 раз в неделю (например, бег, плавание или силовые тренировки).
// Very Active
// Описание: Высокая физическая активность 6–7 раз в неделю (например, интенсивные тренировки или физическая работа).
// Extremely Active
// Описание: Очень высокая физическая активность (например, профессиональные спортсмены или люди с тяжелой физической работой).
const (
	sedentary        = "Sedentary"
	lightlyActive    = "Lightly acive"
	moderatelyActive = "Moderately active"
	veryActive       = "Very active"
	extremelyActive  = "Extremely active"
)

var activityLevels = []string{sedentary, lightlyActive, moderatelyActive, veryActive, extremelyActive}

// Gender
const (
	male   = "Male"
	female = "Female"
)

var genders = []string{male, female}

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
	if !slices.Contains(genders, u.Gender) {
		return fmt.Errorf(
			"gender is error. (%s/%s)",
			male,
			female,
		)
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

	if !slices.Contains(activityLevels, u.ActivityLevel) {
		return fmt.Errorf(
			"ActivityLevel is error. (%s/%s/%s/%s/%s)",
			sedentary,
			lightlyActive,
			moderatelyActive,
			veryActive,
			extremelyActive,
		)
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

// NEW --------------------------------------------------------------------------------------------
// UserRequest представляет весь JSON-объект
type UserRecommendationRequest struct {
	UserID          string         `json:"user_id" validate:"required"`
	UserData        UserData       `json:"user_data" validate:"required"`
	RequestDetails  RequestDetails `json:"request_details" validate:"required"`
	Recommendations string         `json:"recommendations"`
}

// UserData содержит данные пользователя
type UserData struct {
	Profile             Profile             `json:"profile" validate:"required"`
	Goals               Goals               `json:"goals" validate:"required"`
	Preferences         Preferences         `json:"preferences" validate:"required"`
	Lifestyle           Lifestyle           `json:"lifestyle" validate:"required"`
	MedicalRestrictions MedicalRestrictions `json:"medical_restrictions" validate:"required"`
}

// Profile содержит профиль пользователя
type Profile struct {
	Age          int    `json:"age" validate:"required,gt=0,lt=150"`
	Gender       string `json:"gender" validate:"required,oneof=female male"`
	WeightKg     int    `json:"weight_kg" validate:"required,gt=0,lt=300"`
	HeightCm     int    `json:"height_cm" validate:"required,gt=0,lt=300"`
	FitnessLevel string `json:"fitness_level" validate:"required,oneof=beginner intermediate advanced"`
}

// Goals содержит цели пользователя
type Goals struct {
	PrimaryGoal    string `json:"primary_goal" validate:"required,oneof=weight_loss muscle_toning maintenance"`
	SecondaryGoal  string `json:"secondary_goal" validate:"required,oneof=weight_loss muscle_toning maintenance"`
	TargetWeightKg int    `json:"target_weight_kg" validate:"required,gt=0,lt=300"`
	TimeframeWeeks int    `json:"timeframe_weeks" validate:"required,gt=0,lt=52"`
}

// Preferences содержит предпочтения пользователя
type Preferences struct {
	DietType           string   `json:"diet_type" validate:"required,oneof=vegan keto low_carb balanced"`
	Allergies          []string `json:"allergies" validate:"dive,required"`
	PreferredCuisines  []string `json:"preferred_cuisines" validate:"dive,required"`
	WorkoutPreferences []string `json:"workout_preferences" validate:"dive,required"`
}

// Lifestyle содержит информацию об образе жизни пользователя
type Lifestyle struct {
	ActivityLevel           string `json:"activity_level" validate:"required,oneof=sedentary light moderate active very_active"`
	DailyCalorieIntake      int    `json:"daily_calorie_intake" validate:"required,gt=0,lt=5000"`
	WorkoutAvailabilityDays int    `json:"workout_availability_days_per_week" validate:"required,gt=0,lt=8"`
	AverageSleepHours       int    `json:"average_sleep_hours" validate:"required,gt=0,lt=24"`
}

// MedicalRestrictions содержит медицинские ограничения пользователя
type MedicalRestrictions struct {
	HasInjuries       bool     `json:"has_injuries" validate:"required"`
	InjuryDetails     []string `json:"injury_details" validate:"dive,required_with=HasInjuries"`
	ChronicConditions []string `json:"chronic_conditions" validate:"dive,required"`
}

// RequestDetails содержит детали запроса
type RequestDetails struct {
	ServiceType  string `json:"service_type" validate:"required,oneof=fitness_nutrition_recommendations"`
	OutputFormat string `json:"output_format" validate:"required,oneof=daily_plan weekly_plan general_advice"`
	Language     string `json:"language" validate:"required,oneof=ru en"`
}

func (u UserRecommendationRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		// Вывод ошибок валидации
		// for _, err := range err.(validator.ValidationErrors) {
		// 	fmt.Println(err.Namespace(), err.Tag(), err.Param())
		// }
		return err
	}

	return nil
}
