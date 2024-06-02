package model

// LastExerciseSuccess denotes the last successful completion of a given exercise at a given weight
type LastExerciseSuccess struct {
	Lift            Lift   `json:"lift" dynamodbav:"pk"`
	LastSuccessFlag string `json:"-" dynamodbav:"sk"`
	LastSuccessDate string `json:"-" dynamodbav:"date"`
	Weight          int    `json:"workingWeight" dynamodbav:"working_weight"`
}
