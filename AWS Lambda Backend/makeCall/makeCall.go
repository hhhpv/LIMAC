package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Request struct {
	Caller string `json:"caller"`
	Callee string `json:"callee"`
	IP     string `json:"ip"`
}

type Response struct {
	Body   string `json:"body"`
	Status int    `json:"statusCode"`
}

type DBResult struct {
	Caller string `json:"caller"`
	Callee string `json:"callee"`
	IP     string `json:"ip"`
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
		tableName = aws.String(os.Getenv("CALL_TRACKER"))
	)
	callee := request.Callee
	result, err := db_client.GetItem(&dynamodb.GetItemInput{
		TableName: tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"callee": {
				S: aws.String(callee),
			},
		},
	})
	if err != nil {
		return Response{Body: "{\"result\":\"failure\",\"message\":\"dbQueryError\"}", Status: 200}, nil
	}

	if len(result.Item) == 0 {
		entry := &Request{
			Callee: request.Callee,
			Caller: request.Caller,
			IP:     request.IP,
		}
		item, _ := dynamodbattribute.MarshalMap(entry)
		input := &dynamodb.PutItemInput{
			Item:      item,
			TableName: tableName,
		}

		if _, err := db_client.PutItem(input); err != nil {
			return Response{Body: "{\"result\":\"failure\",\"message\":\"dbInsertionErrorOnMakeCall\"}", Status: 200}, nil
		}

		// entry = &Request{
		// 	Callee: request.Caller,
		// 	Caller: request.Callee,
		// 	IP:     request.IP,
		// }
		// item, _ = dynamodbattribute.MarshalMap(entry)
		// input = &dynamodb.PutItemInput{
		// 	Item:      item,
		// 	TableName: tableName,
		// }

		// if _, err = db_client.PutItem(input); err != nil {
		// 	return Response{Body: "{\"result\":\"failure\",\"message\":\"dbInsertionErrorOnMakeCall\"}", Status: 200}, nil
		// }
		return Response{Body: "{\"result\":\"success\",\"message\":\"callPlaced\"}", Status: 200}, nil
	}
	return Response{Body: "{\"result\":\"success\",\"message\":\"calleeBusy\"}", Status: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
