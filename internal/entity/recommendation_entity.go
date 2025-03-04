package entity

import (
	"github.com/go-playground/validator"
)

// UserRequest представляет весь JSON-объект
type UserRecommendationRequest struct {
	UserID          string         `json:"user_id" validate:"required"`
	UserName        string         `json:"user_name" validate:"required"`
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

// Response
type RecommendationResponse struct {
	Recommendations string `json:"recommendations"`
}
