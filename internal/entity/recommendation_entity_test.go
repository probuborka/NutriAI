package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userRecommendationRequestConst = UserRecommendationRequest{
		UserID:   "123456789",
		UserName: "Евгений",
		UserData: UserData{
			Profile: Profile{
				Age:          39,
				Gender:       "male",
				WeightKg:     140,
				HeightCm:     186,
				FitnessLevel: "beginner",
			},
			Goals: Goals{
				PrimaryGoal:    "weight_loss",
				SecondaryGoal:  "muscle_toning",
				TargetWeightKg: 90,
				TimeframeWeeks: 40,
			},
			Preferences: Preferences{
				DietType:           "balanced",
				Allergies:          []string{"орехи", "моллюски"},
				PreferredCuisines:  []string{"средиземноморский", "азиатский"},
				WorkoutPreferences: []string{"йога", "силовая тренировка", "кардио"},
			},
			Lifestyle: Lifestyle{
				ActivityLevel:           "moderate",
				DailyCalorieIntake:      1800,
				WorkoutAvailabilityDays: 4,
				AverageSleepHours:       7,
			},
			MedicalRestrictions: MedicalRestrictions{
				HasInjuries:       true,
				InjuryDetails:     []string{"травма колена"},
				ChronicConditions: []string{"none"},
			},
		},
		RequestDetails: RequestDetails{
			ServiceType:  "fitness_nutrition_recommendations",
			OutputFormat: "weekly_plan",
			Language:     "ru",
		},
	}
)

func TestUserRecommendationRequest_Validate(t *testing.T) {
	type testStruct struct {
		name    string
		request UserRecommendationRequest
		wantErr error
	}
	tests := make([]testStruct, 0)

	//-------------------Valid
	UserRecommendationRequest := userRecommendationRequestConst

	tests = append(tests, testStruct{
		name:    "Valid request",
		request: UserRecommendationRequest,
		wantErr: nil,
	})

	//-------------------Errors
	//--------------------------- UserData
	//--------------------------- UserData.UserID
	//UserID empty
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserID = ""

	tests = append(tests, testStruct{
		name:    "Invalid UserID (empty)",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserID' Error:Field validation for 'UserID' failed on the 'required' tag"),
	})

	//--------------------------- UserData.UserName
	//UserName empty
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserName = ""

	tests = append(tests, testStruct{
		name:    "Invalid UserName (empty)",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserName' Error:Field validation for 'UserName' failed on the 'required' tag"),
	})

	//--------------------------- UserData.Profile
	//--------------------------- UserData.Profile.Age
	//Age empty
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.Age = 0

	tests = append(tests, testStruct{
		name:    "Invalid Age (empty)",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.Age' Error:Field validation for 'Age' failed on the 'required' tag"),
	})

	//Age > 150
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.Age = 151

	tests = append(tests, testStruct{
		name:    "Invalid UserData.Profile.Age (age out of range) > 150",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.Age' Error:Field validation for 'Age' failed on the 'lt' tag"),
	})

	//Age < 1
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.Age = -1

	tests = append(tests, testStruct{
		name:    "Invalid UserData.Profile.Age (age out of range) < 1",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.Age' Error:Field validation for 'Age' failed on the 'gt' tag"),
	})

	//--------------------------- UserData.Profile.Gender
	//Gender empty
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.Gender = ""

	tests = append(tests, testStruct{
		name:    "Invalid Gender (empty)",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.Gender' Error:Field validation for 'Gender' failed on the 'required' tag"),
	})

	//Gender invalid
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.Gender = "invalid_gender"

	tests = append(tests, testStruct{
		name:    "Invalid UserData.Profile.Gender (gender female or male)",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.Gender' Error:Field validation for 'Gender' failed on the 'oneof' tag"),
	})

	//--------------------------- UserData.Profile.WeightKg
	//WeightKg empty
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.WeightKg = 0

	tests = append(tests, testStruct{
		name:    "Invalid WeightKg (empty)",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.WeightKg' Error:Field validation for 'WeightKg' failed on the 'required' tag"),
	})

	//WeightKg > 300
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.WeightKg = 301

	tests = append(tests, testStruct{
		name:    "Invalid UserData.Profile.WeightKg (WeightKg out of range) > 300",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.WeightKg' Error:Field validation for 'WeightKg' failed on the 'lt' tag"),
	})

	//WeightKg < 1
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.WeightKg = -1

	tests = append(tests, testStruct{
		name:    "Invalid UserData.Profile.WeightKg (WeightKg out of range) < 1",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.WeightKg' Error:Field validation for 'WeightKg' failed on the 'gt' tag"),
	})

	//--------------------------- UserData.Profile.HeightCm
	//HeightCm empty
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.HeightCm = 0

	tests = append(tests, testStruct{
		name:    "Invalid HeightCm (empty)",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.HeightCm' Error:Field validation for 'HeightCm' failed on the 'required' tag"),
	})

	//HeightCm > 300
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.HeightCm = 301

	tests = append(tests, testStruct{
		name:    "Invalid UserData.Profile.HeightCm (HeightCm out of range) > 300",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.HeightCm' Error:Field validation for 'HeightCm' failed on the 'lt' tag"),
	})

	//HeightCm < 1
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.HeightCm = -1

	tests = append(tests, testStruct{
		name:    "Invalid UserData.Profile.HeightCm (HeightCm out of range) < 1",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.HeightCm' Error:Field validation for 'HeightCm' failed on the 'gt' tag"),
	})

	//--------------------------- UserData.Profile.FitnessLevel
	//FitnessLevel empty
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.FitnessLevel = ""

	tests = append(tests, testStruct{
		name:    "Invalid FitnessLevel (empty)",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.FitnessLevel' Error:Field validation for 'FitnessLevel' failed on the 'required' tag"),
	})

	//FitnessLevel invalid
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Profile.FitnessLevel = "invalid_FitnessLevel"

	tests = append(tests, testStruct{
		name:    "Invalid UserData.Profile.FitnessLevel (invalid fitnessLevel)",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Profile.FitnessLevel' Error:Field validation for 'FitnessLevel' failed on the 'oneof' tag"),
	})

	//--------------------------- UserData.Goals
	//--------------------------- UserData.Goals.PrimaryGoal
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.UserData.Goals.PrimaryGoal = "invalid_goal"

	tests = append(tests, testStruct{
		name:    "Invalid Goals (invalid primary goal)",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.UserData.Goals.PrimaryGoal' Error:Field validation for 'PrimaryGoal' failed on the 'oneof' tag"),
	})

	//--------------------------- RequestDetails
	//--------------------------- RequestDetails.OutputFormat
	UserRecommendationRequest = userRecommendationRequestConst
	UserRecommendationRequest.RequestDetails.OutputFormat = "invalid_format"

	tests = append(tests, testStruct{
		name:    "Invalid RequestDetails (invalid output format)",
		request: UserRecommendationRequest,
		wantErr: errors.New("Key: 'UserRecommendationRequest.RequestDetails.OutputFormat' Error:Field validation for 'OutputFormat' failed on the 'oneof' tag"),
	})

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
