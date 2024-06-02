package model

// Exercise denotes an attempt at a given physical lift at a given weight
type Exercise struct {
	Lift    Lift   `json:"lift" dynamodbav:"pk"`
	Date    string `json:"-" dynamodbav:"sk"`
	Weight  int    `json:"workingWeight" dynamodbav:"working_weight"`
	Success bool   `json:"-" dynamodbav:"success"`
}
