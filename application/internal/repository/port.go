// Package repository contains adapters for data persistence
package repository

import "jamiekelly/lifts/internal/model"

// Port defining the operations a repository adapter must support
type Port interface {
	CreateExerciseSuccess(model.Exercise)
	CreateExerciseAttempt(model.Exercise)
	ReadPreviousExerciseSuccess(model.Lift) (model.LastExerciseSuccess, error)
	ReadNextWorkoutName() model.WorkoutName
	UpdateNextWorkoutName()
}
