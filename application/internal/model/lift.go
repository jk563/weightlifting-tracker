// Package model holds models
package model

// Lift is an enum for a physical lift
type Lift string

// Acceptable types of lift
const (
	Squat         Lift = "Squat"
	BarbellRow    Lift = "Barbell Row"
	Deadlift      Lift = "Deadlift"
	OverheadPress Lift = "Overhead Press"
	BenchPress    Lift = "Bench Press"
	BicepCurl     Lift = "Bicep Curl"
)

// Minimum returns the default (minimum) working weight for the given lift
func (l Lift) Minimum() int {
	switch l {
	case Squat, OverheadPress, BenchPress, BicepCurl:
		return 15000
	case BarbellRow, Deadlift:
		return 25000
	default:
		return 0
	}
}
