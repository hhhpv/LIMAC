package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"golang.org/x/crypto/bcrypt"
)

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Body   string `json:"body"`
	Status int    `json:"statusCode"`
}

var db_client *dynamodb.DynamoDB

func init() {
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		db_client = dynamodb.New(session) // Create DynamoDB client
	}
}

func HandleRequest(ctx context.Context, request Request) (Response, error) {
	var (
		tableName = aws.String(os.Getenv("USERS"))
	)
	username := request.Username

	result, err := db_client.GetItem(&dynamodb.GetItemInput{
		TableName: tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})
	if err != nil {
		return Response{Body: "{\"result\":\"failure\",\"message\":\"dbQueryError\"}", Status: 200}, nil
	}

	if len(result.Item) == 0 {
		entry := &Request{
			Username: request.Username,
			Password: hashAndSalt([]byte(request.Password), 4),
		}
		item, _ := dynamodbattribute.MarshalMap(entry)
		input := &dynamodb.PutItemInput{
			Item:      item,
			TableName: tableName,
		}

		if _, err := db_client.PutItem(input); err != nil {
			return Response{Body: "{\"result\":\"failure\",\"message\":\"dbInsertionErrorOnSignUp\"}", Status: 200}, nil
		}
	} else {
		return Response{Body: "{\"result\":\"failure\",\"message\":\"usernameAlreadyInUse\"}", Status: 200}, nil
	}
	return Response{Body: "{\"result\":\"success\",\"message\":\"userCreated\"}", Status: 200}, nil
}

func hashAndSalt(pwd []byte, cost int) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, cost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func main() {
	lambda.Start(HandleRequest)
}
