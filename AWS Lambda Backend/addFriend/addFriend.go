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
	Friend   string `json:"friend"`
}

type Response struct {
	Body   string `json:"body"`
	Status int    `json:"statusCode"`
}

type DBResult struct {
	Username string   `json:"username"`
	Friends  []string `json:"friends"`
}

type FriendListUpdate struct {
	Friends []string `json:":r"`
}

type FriendListNew struct {
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
		tableName_friends = aws.String(os.Getenv("FRIENDS"))
	)
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
	if len(result.Item) >= 1 {
		var db_result DBResult
		decode_err := dynamodbattribute.UnmarshalMap(result.Item, &db_result)
		if decode_err != nil {
			return Response{Body: "{\"result\":\"failure\",\"message\":\"decodeError\"}", Status: 200}, nil
		}
		var friendList []string
		for _, item := range db_result.Friends {
			friendList = append(friendList, item)
		}
		friendList = append(friendList, request.Friend)
		update, marshal_err := dynamodbattribute.MarshalMap(FriendListUpdate{
			Friends: friendList,
		})
		if marshal_err != nil {
			fmt.Println(marshal_err)
		}
		input := &dynamodb.UpdateItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"username": {
					S: aws.String(username),
				}},
			TableName:                 tableName_friends,
			UpdateExpression:          aws.String("SET friends= :r"),
			ExpressionAttributeValues: update,
			ReturnValues:              aws.String("UPDATED_NEW"),
		}
		_, updateErr := db_client.UpdateItem(input)
		if updateErr != nil {
			return Response{Body: "{\"result\":\"failure\",\"message\":\"updateError\"}", Status: 200}, nil
		}
		return Response{Body: "{\"result\":\"success\",\"message\":\"addedFriend\"}", Status: 200}, nil
	}
	var friendList []string
	friendList = append(friendList, request.Friend)
	entry := &FriendListNew{
		Username: request.Username,
		Friends:  friendList,
	}
	item, _ := dynamodbattribute.MarshalMap(entry)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: tableName_friends,
	}
	if _, err := db_client.PutItem(input); err != nil {
		return Response{Body: "{\"result\":\"failure\",\"message\":\"dbInsertionErrorFriendMap\"}", Status: 200}, nil
	}
	return Response{Body: "{\"result\":\"success\",\"message\":\"addedFriend\"}", Status: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
