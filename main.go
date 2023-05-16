package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, payload events.CloudWatchEvent) (string, error) {
	p()
	return string(payload.Region), nil
}

func main() {
	lambda.Start(HandleRequest)
}