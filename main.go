package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"context"
)

type event struct {
	name string
}

func HandleRequest(ctx context.Context, name event)(string, error) {
	p()
	return string(name.name), nil
}

func main() {
	lambda.Start(HandleRequest)
}