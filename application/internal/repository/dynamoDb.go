package repository

import (
	"fmt"
	"jamiekelly/lifts/internal/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rs/zerolog/log"
)

// DynamoDb repository adaptor
type DynamoDb struct {
	client    *dynamodb.DynamoDB
	tableName string
}

// DynamoDbAdapter initialises and returns a DynamoDB object to use for persistence
func DynamoDbAdapter() DynamoDb {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ddbClient := dynamodb.New(session)
	return DynamoDb{
		client:    ddbClient,
		tableName: "lifts",
	}
}

// CreateExerciseSuccess stores an exercise success
func (ddb DynamoDb) CreateExerciseSuccess(exercise model.Exercise) {
	lastExerciseSuccess := model.LastExerciseSuccess{
		Lift:            exercise.Lift,
		LastSuccessFlag: "lastSuccess",
		LastSuccessDate: exercise.Date,
		Weight:          exercise.Weight,
	}

	newItem, err := dynamodbattribute.MarshalMap(lastExerciseSuccess)
	if err != nil {
		log.Fatal().Msgf("Got error marshalling new exercise item: %s", err)
	}

	newInput := &dynamodb.PutItemInput{
		Item:      newItem,
		TableName: aws.String(ddb.tableName),
	}

	_, err = ddb.client.PutItem(newInput)
	if err != nil {
		log.Fatal().Msgf("Got error calling PutItem: %s", err)
	}
}

// CreateExerciseAttempt stores an exercise attempt
func (ddb DynamoDb) CreateExerciseAttempt(exercise model.Exercise) {
	newItem, err := dynamodbattribute.MarshalMap(exercise)
	if err != nil {
		log.Fatal().Msgf("Got error marshalling new exercise item: %s", err)
	}

	newInput := &dynamodb.PutItemInput{
		Item:      newItem,
		TableName: aws.String(ddb.tableName),
	}

	_, err = ddb.client.PutItem(newInput)
	if err != nil {
		log.Fatal().Msgf("Got error calling PutItem: %s", err)
	}
}

// ReadPreviousExerciseSuccess retrieves the last successful workout for an exercise from storage
func (ddb DynamoDb) ReadPreviousExerciseSuccess(lift model.Lift) (model.LastExerciseSuccess, error) {
	result, err := ddb.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(ddb.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(string(lift)),
			},
			"sk": {
				S: aws.String("lastSuccess"),
			},
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Error getting last successful lift:, %v", err))
	}

	previousExerciseSuccess := model.LastExerciseSuccess{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &previousExerciseSuccess)
	if err != nil {
		log.Error().Err(err).Str("exercise", string(lift)).Msg("Error unmarshalling")
	}
	return previousExerciseSuccess, nil
}

// ReadNextWorkoutName retrieves the next workout from storage
func (ddb DynamoDb) ReadNextWorkoutName() model.WorkoutName {
	result, err := ddb.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(ddb.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String("workout"),
			},
			"sk": {
				S: aws.String("next"),
			},
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	type dynamodbWorkoutNameRecord struct {
		WorkoutName model.WorkoutName `dynamodbav:"exercises"`
	}
	workoutNameRecord := dynamodbWorkoutNameRecord{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &workoutNameRecord)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return workoutNameRecord.WorkoutName
}

// UpdateNextWorkoutName updates the next workout name to the next workout in the cycle
func (ddb DynamoDb) UpdateNextWorkoutName() {
	workoutName := ddb.ReadNextWorkoutName()

	type dynamodbNextWorkoutRecord struct {
		WorkoutName model.WorkoutName `dynamodbav:"exercises"`
		PK          string            `dynamodbav:"pk"`
		SK          string            `dynamodbav:"sk"`
	}

	nextWorkoutNameRecord := dynamodbNextWorkoutRecord{
		PK: "workout",
		SK: "next",
	}

	if workoutName == model.WorkoutA {
		nextWorkoutNameRecord.WorkoutName = model.WorkoutB
	} else {
		nextWorkoutNameRecord.WorkoutName = model.WorkoutA
	}

	newItem, err := dynamodbattribute.MarshalMap(nextWorkoutNameRecord)
	if err != nil {
		log.Fatal().Msgf("Got error marshalling new workout item: %s", err)
	}

	newInput := &dynamodb.PutItemInput{
		Item:      newItem,
		TableName: aws.String(ddb.tableName),
	}

	_, err = ddb.client.PutItem(newInput)
	if err != nil {
		log.Fatal().Msgf("Got error calling PutItem: %s", err)
	}
}
