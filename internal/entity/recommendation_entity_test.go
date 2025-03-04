package entity

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		user    UserNutritionAndFitnessProfile
		wantErr error
	}{
		{
			name: "Valid user",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                25,
				Gender:             "male",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: nil,
		},
		{
			name: "Empty UserID",
			user: UserNutritionAndFitnessProfile{
				UserID:             "",
				Age:                25,
				Gender:             "male",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("user_id is required"),
		},
		{
			name: "UserID too short",
			user: UserNutritionAndFitnessProfile{
				UserID:             "ab",
				Age:                25,
				Gender:             "male",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("user_id must be between 3 and 50 characters"),
		},
		{
			name: "UserID too long",
			user: UserNutritionAndFitnessProfile{
				UserID:             "thisisaverylonguseridthatiswaybeyondtheallowedlimitof50characters",
				Age:                25,
				Gender:             "male",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("user_id must be between 3 and 50 characters"),
		},
		{
			name: "Invalid age (too young)",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                17,
				Gender:             "male",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("age must be between 18 and 100"),
		},
		{
			name: "Invalid age (too old)",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                101,
				Gender:             "male",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("age must be between 18 and 100"),
		},
		{
			name: "Empty gender",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                25,
				Gender:             "",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("gender is required"),
		},
		{
			name: "Invalid gender",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                25,
				Gender:             "other",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("gender is error. (male or female)"),
		},
		{
			name: "Empty height",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                25,
				Gender:             "male",
				Height:             0,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("height is required"),
		},
		{
			name: "Invalid height (too short)",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                25,
				Gender:             "male",
				Height:             40,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("height must be between 50 and 250 cm"),
		},
		{
			name: "Invalid height (too tall)",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                25,
				Gender:             "male",
				Height:             260,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("height must be between 50 and 250 cm"),
		},
		{
			name: "Empty currentWeight",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                25,
				Gender:             "male",
				Height:             180,
				CurrentWeight:      0,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("currentWeight is required"),
		},
		{
			name: "Empty goalWeight",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                25,
				Gender:             "male",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         0,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("goalWeight is required"),
		},
		{
			name: "Empty activityLevel",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                25,
				Gender:             "male",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "",
				DietaryPreferences: "vegan",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("activityLevel is required"),
		},
		{
			name: "Empty dietaryPreferences",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                25,
				Gender:             "male",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "",
				TrainingGoals:      "strength",
			},
			wantErr: errors.New("dietaryPreferences is required"),
		},
		{
			name: "Empty trainingGoals",
			user: UserNutritionAndFitnessProfile{
				UserID:             "user123",
				Age:                25,
				Gender:             "male",
				Height:             180,
				CurrentWeight:      70,
				GoalWeight:         65,
				ActivityLevel:      "active",
				DietaryPreferences: "vegan",
				TrainingGoals:      "",
			},
			wantErr: errors.New("trainingGoals is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil && tt.wantErr == nil) ||
				(err == nil && tt.wantErr != nil) ||
				(err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
