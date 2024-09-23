// Package app contains core business logic
package app

import (
	"jamiekelly/lifts/internal/model"
	"jamiekelly/lifts/internal/repository"
	"time"

	"github.com/rs/zerolog/log"
)

// App is the main application API
type App struct {
	Repo repository.Port
}

var workoutA = []model.Lift{"Bench Press", "Barbell Row", "Squat"}
var workoutB = []model.Lift{"Overhead Press", "Bicep Curl", "Deadlift"}

// GetNextWorkout retrieves the next workout to do
func (app App) GetNextWorkout() model.Workout {
	log.Debug().Msg("Retrieving next workout name")
	workoutName := app.Repo.ReadNextWorkoutName()
	log.Info().Str("Name", string(workoutName)).Msg("Next one")

	log.Debug().Msg("Working out next lifts to attempt")
	exercises := []model.Exercise{}
	if workoutName == "workoutA" {
		log.Info().Str("workout", "workoutA").Msg("Next workout")
		exercises = []model.Exercise{
			{Lift: model.BenchPress},
			{Lift: model.BarbellRow},
			{Lift: model.Squat},
		}
	} else {
		log.Info().Str("workout", "workoutB").Msg("Next workout")
		exercises = []model.Exercise{
			{Lift: model.OverheadPress},
			{Lift: model.BicepCurl},
			{Lift: model.Deadlift},
		}
	}

	for i, exercise := range exercises {
		exerciseLogger := log.With().Str("exercise", string(exercise.Lift)).Logger()
		exerciseLogger.Debug().Msg("Get previous entry for exercise")

		previousExercise, err := app.Repo.ReadPreviousExerciseSuccess(exercise.Lift)
		if err != nil {
			exerciseLogger.Debug().Msg("Error getting previous, using minimum")
			exercises[i].Weight = exercise.Lift.Minimum()
		} else {
			previousSuccessDate, err := time.Parse("20060102", previousExercise.LastSuccessDate)
			if err != nil {
				exerciseLogger.Debug().Msg("Error parsing previous success date, using minimum")
				exercises[i].Weight = exercise.Lift.Minimum()
			} else {
				currentTime := time.Now()
				nextWeight := previousExercise.Weight + 2500
				daysSinceLastSuccess := int(currentTime.Sub(previousSuccessDate).Hours() / 24)
				if daysSinceLastSuccess <= 14 {
					exerciseLogger.Debug().Msg("Exercise success recently, increasing weight from last succss")
					exercises[i].Weight = nextWeight
				} else if daysSinceLastSuccess >= 90 {
					exerciseLogger.Info().Msg("Not done in 90+ days, starting at minimum")
					exercises[i].Weight = exercise.Lift.Minimum()
				} else {
					numberOfReductions := daysSinceLastSuccess / 14
					reducedWeight := ((((nextWeight / 10) * (10 - numberOfReductions)) / 2500) * 2500)
					if reducedWeight <= exercise.Lift.Minimum() {
						exercises[i].Weight = exercise.Lift.Minimum()
						exerciseLogger.Debug().Msgf("Exercise not completed in %v days, lowering weight to minimumv", daysSinceLastSuccess)
					} else {
						exercises[i].Weight = reducedWeight
						exerciseLogger.Debug().Msgf("Exercise not completed in %v days, lowering weight %v times to %v", daysSinceLastSuccess, numberOfReductions, reducedWeight)
					}
				}
			}
		}
	}

	return model.Workout{
		WorkoutName: workoutName,
		Exercises:   exercises,
	}
}

// CompleteWorkout persists a finished workout
func (app App) CompleteWorkout(success bool) {
	currentWorkout := app.GetNextWorkout()

	for _, exercise := range currentWorkout.Exercises {
		currentTime := time.Now()
		exercise.Date = currentTime.Format("20060102")
		exercise.Success = success

		app.Repo.CreateExerciseAttempt(exercise)

		if success {
			app.Repo.CreateExerciseSuccess(exercise)
		}
	}

	app.Repo.UpdateNextWorkoutName()
}
