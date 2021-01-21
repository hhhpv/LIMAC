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

type Intermediate struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

type void struct{}

func HandleRequest(ctx context.Context, request Request) (Response, error) {
	var (
		tableName_friends = aws.String(os.Getenv("FRIENDS"))
		tableName_users   = aws.String(os.Getenv("USERS"))
		none              void
	)
	set_friends := make(map[string]void)
	set_members := make(map[string]void)
	username := request.Username
	result, err := db_client.GetItem(&dynamodb.GetItemInput{
		TableName: tableName_friends,
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

	if len(db_result.Friends) >= 1 {
		for _, item := range db_result.Friends {
			set_friends[item] = none
		}
	}

	scan_input := &dynamodb.ScanInput{
		TableName: tableName_users,
	}
	scan_result, _ := db_client.Scan(scan_input)
	var decode_item Intermediate
	for _, item := range scan_result.Items {
		decode_err = dynamodbattribute.UnmarshalMap(item, &decode_item)
		_, exists := set_friends[decode_item.Username]
		if !exists {
			set_members[decode_item.Username] = none
		}
	}
	list := "["
	for item := range set_members { // Loop
		list += "\"" + item + "\"" + ","
	}
	sz := len(list)
	list = list[:sz-1] + "]"
	return Response{Body: "{\"result\":\"success\",\"message\":" + list + "}", Status: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
