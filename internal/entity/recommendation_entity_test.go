package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRecommendationRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request UserRecommendationRequest
		wantErr error
	}{
		{
			name: "Valid request",
			request: UserRecommendationRequest{
				UserID: "12345",
				UserData: UserData{
					Profile: Profile{
						Age:          30,
						Gender:       "female",
						WeightKg:     70,
						HeightCm:     165,
						FitnessLevel: "intermediate",
					},
					Goals: Goals{
						PrimaryGoal:    "weight_loss",
						SecondaryGoal:  "muscle_toning",
						TargetWeightKg: 65,
						TimeframeWeeks: 12,
					},
					Preferences: Preferences{
						DietType:           "balanced",
						Allergies:          []string{"nuts"},
						PreferredCuisines:  []string{"mediterranean"},
						WorkoutPreferences: []string{"yoga"},
					},
					Lifestyle: Lifestyle{
						ActivityLevel:           "moderate",
						DailyCalorieIntake:      1800,
						WorkoutAvailabilityDays: 4,
						AverageSleepHours:       7,
					},
					MedicalRestrictions: MedicalRestrictions{
						HasInjuries:       true,
						InjuryDetails:     []string{"lower_back_pain"},
						ChronicConditions: []string{"none"},
					},
				},
				RequestDetails: RequestDetails{
					ServiceType:  "fitness_nutrition_recommendations",
					OutputFormat: "weekly_plan",
					Language:     "ru",
				},
			},
			wantErr: nil,
		},
		{
			name: "Invalid UserID (empty)",
			request: UserRecommendationRequest{
				UserID: "",
				UserData: UserData{
					Profile: Profile{
						Age:          30,
						Gender:       "female",
						WeightKg:     70,
						HeightCm:     165,
						FitnessLevel: "intermediate",
					},
					Goals: Goals{
						PrimaryGoal:    "weight_loss",
						SecondaryGoal:  "muscle_toning",
						TargetWeightKg: 65,
						TimeframeWeeks: 12,
					},
					Preferences: Preferences{
						DietType:           "balanced",
						Allergies:          []string{"nuts"},
						PreferredCuisines:  []string{"mediterranean"},
						WorkoutPreferences: []string{"yoga"},
					},
					Lifestyle: Lifestyle{
						ActivityLevel:           "moderate",
						DailyCalorieIntake:      1800,
						WorkoutAvailabilityDays: 4,
						AverageSleepHours:       7,
					},
					MedicalRestrictions: MedicalRestrictions{
						HasInjuries:       true,
						InjuryDetails:     []string{"lower_back_pain"},
						ChronicConditions: []string{"none"},
					},
				},
				RequestDetails: RequestDetails{
					ServiceType:  "fitness_nutrition_recommendations",
					OutputFormat: "weekly_plan",
					Language:     "ru",
				},
			},
			wantErr: errors.New("Key: 'UserRecommendationRequest.UserID' Error:Field validation for 'UserID' failed on the 'required' tag"),
		},
		{
			name: "Invalid Profile (age out of range)",
			request: UserRecommendationRequest{
				UserID: "12345",
				UserData: UserData{
					Profile: Profile{
						Age:          200,
						Gender:       "female",
						WeightKg:     70,
						HeightCm:     165,
						FitnessLevel: "intermediate",
					},
					Goals: Goals{
						PrimaryGoal:    "weight_loss",
						SecondaryGoal:  "muscle_toning",
						TargetWeightKg: 65,
						TimeframeWeeks: 12,
					},
					Preferences: Preferences{
						DietType:           "balanced",
						Allergies:          []string{"nuts"},
						PreferredCuisines:  []string{"mediterranean"},
						WorkoutPreferences: []string{"yoga"},
					},
					Lifestyle: Lifestyle{
						ActivityLevel:           "moderate",
						DailyCalorieIntake:      1800,
						WorkoutAvailabilityDays: 4,
						AverageSleepHours:       7,
					},
					MedicalRestrictions: MedicalRestrictions{
						HasInjuries:       true,
						InjuryDetails:     []string{"lower_back_pain"},
						ChronicConditions: []string{"none"},
					},
				},
				RequestDetails: RequestDetails{
					ServiceType:  "fitness_nutrition_recommendations",
					OutputFormat: "weekly_plan",
					Language:     "ru",
				},
			},
			wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.Age' Error:Field validation for 'Age' failed on the 'lt' tag"),
		},
		{
			name: "Invalid Goals (invalid primary goal)",
			request: UserRecommendationRequest{
				UserID: "12345",
				UserData: UserData{
					Profile: Profile{
						Age:          30,
						Gender:       "female",
						WeightKg:     70,
						HeightCm:     165,
						FitnessLevel: "intermediate",
					},
					Goals: Goals{
						PrimaryGoal:    "invalid_goal",
						SecondaryGoal:  "muscle_toning",
						TargetWeightKg: 65,
						TimeframeWeeks: 12,
					},
					Preferences: Preferences{
						DietType:           "balanced",
						Allergies:          []string{"nuts"},
						PreferredCuisines:  []string{"mediterranean"},
						WorkoutPreferences: []string{"yoga"},
					},
					Lifestyle: Lifestyle{
						ActivityLevel:           "moderate",
						DailyCalorieIntake:      1800,
						WorkoutAvailabilityDays: 4,
						AverageSleepHours:       7,
					},
					MedicalRestrictions: MedicalRestrictions{
						HasInjuries:       true,
						InjuryDetails:     []string{"lower_back_pain"},
						ChronicConditions: []string{"none"},
					},
				},
				RequestDetails: RequestDetails{
					ServiceType:  "fitness_nutrition_recommendations",
					OutputFormat: "weekly_plan",
					Language:     "ru",
				},
			},
			wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Goals.PrimaryGoal' Error:Field validation for 'PrimaryGoal' failed on the 'oneof' tag"),
		},
		{
			name: "Invalid RequestDetails (invalid output format)",
			request: UserRecommendationRequest{
				UserID: "12345",
				UserData: UserData{
					Profile: Profile{
						Age:          30,
						Gender:       "female",
						WeightKg:     70,
						HeightCm:     165,
						FitnessLevel: "intermediate",
					},
					Goals: Goals{
						PrimaryGoal:    "weight_loss",
						SecondaryGoal:  "muscle_toning",
						TargetWeightKg: 65,
						TimeframeWeeks: 12,
					},
					Preferences: Preferences{
						DietType:           "balanced",
						Allergies:          []string{"nuts"},
						PreferredCuisines:  []string{"mediterranean"},
						WorkoutPreferences: []string{"yoga"},
					},
					Lifestyle: Lifestyle{
						ActivityLevel:           "moderate",
						DailyCalorieIntake:      1800,
						WorkoutAvailabilityDays: 4,
						AverageSleepHours:       7,
					},
					MedicalRestrictions: MedicalRestrictions{
						HasInjuries:       true,
						InjuryDetails:     []string{"lower_back_pain"},
						ChronicConditions: []string{"none"},
					},
				},
				RequestDetails: RequestDetails{
					ServiceType:  "fitness_nutrition_recommendations",
					OutputFormat: "invalid_format",
					Language:     "ru",
				},
			},
			wantErr: errors.New("Key: 'UserRecommendationRequest.RequestDetails.OutputFormat' Error:Field validation for 'OutputFormat' failed on the 'oneof' tag"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}

// import (
// 	"errors"
// 	"testing"
// )

// func TestValidate(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		user    UserNutritionAndFitnessProfile
// 		wantErr error
// 	}{
// 		{
// 			name: "Valid user",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: nil,
// 		},
// 		{
// 			name: "Empty UserID",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("user_id is required"),
// 		},
// 		{
// 			name: "UserID too short",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "ab",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("user_id must be between 3 and 50 characters"),
// 		},
// 		{
// 			name: "UserID too long",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "thisisaverylonguseridthatiswaybeyondtheallowedlimitof50characters",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("user_id must be between 3 and 50 characters"),
// 		},
// 		{
// 			name: "Invalid age (too young)",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                17,
// 				Gender:             "male",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("age must be between 18 and 100"),
// 		},
// 		{
// 			name: "Invalid age (too old)",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                101,
// 				Gender:             "male",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("age must be between 18 and 100"),
// 		},
// 		{
// 			name: "Empty gender",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                25,
// 				Gender:             "",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("gender is required"),
// 		},
// 		{
// 			name: "Invalid gender",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                25,
// 				Gender:             "other",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("gender is error. (male or female)"),
// 		},
// 		{
// 			name: "Empty height",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             0,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("height is required"),
// 		},
// 		{
// 			name: "Invalid height (too short)",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             40,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("height must be between 50 and 250 cm"),
// 		},
// 		{
// 			name: "Invalid height (too tall)",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             260,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("height must be between 50 and 250 cm"),
// 		},
// 		{
// 			name: "Empty currentWeight",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             180,
// 				CurrentWeight:      0,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("currentWeight is required"),
// 		},
// 		{
// 			name: "Empty goalWeight",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         0,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("goalWeight is required"),
// 		},
// 		{
// 			name: "Empty activityLevel",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("activityLevel is required"),
// 		},
// 		{
// 			name: "Empty dietaryPreferences",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "",
// 				TrainingGoals:      "strength",
// 			},
// 			wantErr: errors.New("dietaryPreferences is required"),
// 		},
// 		{
// 			name: "Empty trainingGoals",
// 			user: UserNutritionAndFitnessProfile{
// 				UserID:             "user123",
// 				Age:                25,
// 				Gender:             "male",
// 				Height:             180,
// 				CurrentWeight:      70,
// 				GoalWeight:         65,
// 				ActivityLevel:      "active",
// 				DietaryPreferences: "vegan",
// 				TrainingGoals:      "",
// 			},
// 			wantErr: errors.New("trainingGoals is required"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := tt.user.Validate()
// 			if (err != nil && tt.wantErr == nil) ||
// 				(err == nil && tt.wantErr != nil) ||
// 				(err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
// 				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
