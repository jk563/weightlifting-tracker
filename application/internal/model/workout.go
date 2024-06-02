package model

// Workout denotes a workout to complete
type Workout struct {
	WorkoutName WorkoutName `json:"workoutName" dynamodbav:"exercises"`
	Exercises   []Exercise  `json:"exercises" dynamodbav:"-"`
}
