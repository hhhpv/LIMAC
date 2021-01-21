package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Ques string `json:"ques"`
}

type Response struct {
	Body   string `json:"body"`
	Status int    `json:"statusCode"`
}

func HandleRequest(ctx context.Context, request Request) (Response, error) {
	return Response{Body: "{\"username\":\"hit\",\"password\":\"hit123\",\"ip\":\"13.56.194.175\",\"port\":\"3478\"}", Status: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
