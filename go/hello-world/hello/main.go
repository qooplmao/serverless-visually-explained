package main

import (
	"log"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(context context.Context) (string, error) {
	log.Print("And running... running...")

	return "Hello World!", nil
}

func main() {
	lambda.Start(Handler)
}