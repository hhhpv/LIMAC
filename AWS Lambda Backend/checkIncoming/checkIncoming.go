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
	Caller string `json:"caller"`
	Callee string `json:"callee"`
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
	// caller := request.Caller
	callee := request.Callee
	result, err := db_client.GetItem(&dynamodb.GetItemInput{
		TableName: tableName,
		Key: map[string]*dynamodb.AttributeValue{
			// "caller": {
			// 	S: aws.String(caller),
			// },
			"callee": {
				S: aws.String(callee),
			},
		},
	})
	if err != nil {
		return Response{Body: "{\"result\":\"failure\",\"message\":\"dbQueryError\"}", Status: 200}, nil
	}

	if len(result.Item) == 0 {
		return Response{Body: "{\"result\":\"success\",\"message\":\"NoIncomingCall\"}", Status: 200}, nil
	}
	var db_result DBResult
	decode_err := dynamodbattribute.UnmarshalMap(result.Item, &db_result)
	if decode_err != nil {
		return Response{Body: "{\"result\":\"failure\",\"message\":\"decodeError\"}", Status: 200}, nil
	}
	return Response{Body: "{\"result\":\"success\",\"message\":\"CallIncoming\",\"callerIP\":\"" + db_result.IP + "\",\"caller\":\"" + db_result.Caller + "\"}", Status: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
