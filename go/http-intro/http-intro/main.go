package main

import (
	"log"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.QueryStringParameters["name"]
	log.Printf("Name: %s", name)
	body := "Hello World!"

	if 0 < len(name) {
		body = fmt.Sprintf("Hi %s", name)
	}

	fmt.Print(body)
	log.Print(body)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: body,
	}, nil
}

func main() {
	lambda.Start(Handler)
}