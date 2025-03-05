package entity

import (
	"github.com/go-playground/validator"
)

// UserRecommendationRequest представляет модель пользователя
// @Description Информация о пользователе
type UserRecommendationRequest struct {
	UserID          string         `json:"user_id" example:"123456789" validate:"required"`
	UserName        string         `json:"user_name" example:"Евгений" validate:"required"`
	UserData        UserData       `json:"user_data" validate:"required"`
	RequestDetails  RequestDetails `json:"request_details" validate:"required"`
	Recommendations string         `json:"recommendations"`
}

// UserData содержит данные пользователя
// @Description содержит данные пользователя
type UserData struct {
	Profile             Profile             `json:"profile" validate:"required"`
	Goals               Goals               `json:"goals" validate:"required"`
	Preferences         Preferences         `json:"preferences" validate:"required"`
	Lifestyle           Lifestyle           `json:"lifestyle" validate:"required"`
	MedicalRestrictions MedicalRestrictions `json:"medical_restrictions" validate:"required"`
}

// Profile содержит профиль пользователя
// @Description содержит профиль пользователя
type Profile struct {
	Age          int    `json:"age" example:"39" validate:"required,gt=0,lt=150"`
	Gender       string `json:"gender" example:"male" validate:"required,oneof=female male"`
	WeightKg     int    `json:"weight_kg" example:"140" validate:"required,gt=0,lt=300"`
	HeightCm     int    `json:"height_cm" example:"186" validate:"required,gt=0,lt=300"`
	FitnessLevel string `json:"fitness_level" example:"beginner" validate:"required,oneof=beginner intermediate advanced"`
}

// Goals содержит цели пользователя
// @Description содержит цели пользователя
type Goals struct {
	PrimaryGoal    string `json:"primary_goal" example:"weight_loss" validate:"required,oneof=weight_loss muscle_toning maintenance"`
	SecondaryGoal  string `json:"secondary_goal" example:"muscle_toning" validate:"required,oneof=weight_loss muscle_toning maintenance"`
	TargetWeightKg int    `json:"target_weight_kg" example:"90" validate:"required,gt=0,lt=300"`
	TimeframeWeeks int    `json:"timeframe_weeks" example:"40" validate:"required,gt=0,lt=52"`
}

// Preferences содержит предпочтения пользователя
// @Description содержит предпочтения пользователя
type Preferences struct {
	DietType           string   `json:"diet_type" example:"balanced" validate:"required,oneof=vegan keto low_carb balanced"`
	Allergies          []string `json:"allergies" example:"орехи,моллюски" validate:"dive,required"`
	PreferredCuisines  []string `json:"preferred_cuisines" example:"средиземноморский,азиатский" validate:"dive,required"`
	WorkoutPreferences []string `json:"workout_preferences" example:"силовая тренировка,кардио" validate:"dive,required"`
}

// Lifestyle содержит информацию об образе жизни пользователя
// @Description содержит информацию об образе жизни пользователя
type Lifestyle struct {
	ActivityLevel           string `json:"activity_level" example:"moderate" validate:"required,oneof=sedentary light moderate active very_active"`
	DailyCalorieIntake      int    `json:"daily_calorie_intake" example:"1800" validate:"required,gt=0,lt=5000"`
	WorkoutAvailabilityDays int    `json:"workout_availability_days_per_week" example:"4" validate:"required,gt=0,lt=8"`
	AverageSleepHours       int    `json:"average_sleep_hours" example:"7" validate:"required,gt=0,lt=24"`
}

// MedicalRestrictions содержит медицинские ограничения пользователя
// @Description содержит медицинские ограничения пользователя
type MedicalRestrictions struct {
	HasInjuries       bool     `json:"has_injuries" example:"true"  validate:"required"`
	InjuryDetails     []string `json:"injury_details" example:"травма колена"  validate:"dive,required_with=HasInjuries"`
	ChronicConditions []string `json:"chronic_conditions" example:"none"  validate:"dive,required"`
}

// RequestDetails содержит детали запроса
// @Description содержит детали запроса
type RequestDetails struct {
	ServiceType  string `json:"service_type" example:"fitness_nutrition_recommendations" validate:"required,oneof=fitness_nutrition_recommendations"`
	OutputFormat string `json:"output_format" example:"weekly_plan" validate:"required,oneof=daily_plan weekly_plan general_advice"`
	Language     string `json:"language" example:"ru" validate:"required,oneof=ru en"`
}

func (u UserRecommendationRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}

	return nil
}

// Response
type RecommendationResponse struct {
	Recommendations string `json:"recommendations"`
}
