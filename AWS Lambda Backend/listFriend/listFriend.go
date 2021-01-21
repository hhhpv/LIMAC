package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Request struct {
	Username string `json:"username"`
}

type Response struct {
	Body   string `json:"body"`
	Status int    `json:"statusCode"`
}

type DBResult struct {
	Username string   `json:"username"`
	Friends  []string `json:"friends"`
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
		tableName = aws.String(os.Getenv("FRIENDS"))
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
	var db_result DBResult
	decode_err := dynamodbattribute.UnmarshalMap(result.Item, &db_result)
	if decode_err != nil {
		return Response{Body: "{\"result\":\"failure\",\"message\":\"decodeError\"}", Status: 200}, nil
	}
	friendList := "["
	if len(db_result.Friends) >= 1 {
		for _, item := range db_result.Friends {
			friendList += "\"" + item + "\"" + ","
		}
		sz := len(friendList)
		friendList = friendList[:sz-1] + "]"
		return Response{Body: "{\"result\":\"success\",\"message\":" + friendList + "}", Status: 200}, nil
	}
	friendList = friendList + "]"
	return Response{Body: "{\"result\":\"success\",\"message\":" + friendList + "}", Status: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
