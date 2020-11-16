package main

import (
	"log"
	"context"
	"strings"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Name string `json:"name"`
}

type Response struct {
	Lowercase string `json:"lowercase"`
	Uppercase string `json:"uppercase"`
}

func Handler(context context.Context, event Event) (Response, error) {
	name := event.Name
	log.Printf("Name: %s", name)

	if 0 == len(name) {
		name = "None set"
		log.Print("No name, name set to \"None Set\"")
		log.Printf("Name: %s", name)
	}

	return Response{
		Lowercase: strings.ToLower(name),
		Uppercase: strings.ToUpper(name),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
