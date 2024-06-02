package driver

import (
	"context"
	"errors"
	"jamiekelly/lifts/internal/app"
	"jamiekelly/lifts/internal/model"
	"jamiekelly/lifts/internal/repository"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

// Lambda defines the Lambda driver adapter
type Lambda struct{}

var application app.App

// Init sets up the core application adapters
func (driver Lambda) Init() {
	application = app.App{
		Repo: repository.DynamoDbAdapter(),
	}
}

// Run starts the Lambda runtime for handling requests
func (driver Lambda) Run() {
	lambda.Start(HandleRequest)
}

type myEvent struct{}

/*
HandleRequest handles Lambda invocation
*/
func HandleRequest(_ context.Context, _ myEvent) (model.Workout, error) {
	requestType := os.Getenv("RequestType")

	if requestType == "GetWorkout" {
		return application.GetNextWorkout(), nil
	}

	if requestType == "CompleteWorkout" {
		application.CompleteWorkout(true)
		return model.Workout{}, nil
	}

	return model.Workout{}, errors.New("Not sure what type of request to execute")
}
