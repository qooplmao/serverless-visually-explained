package main

import (
	"encoding/json"
	"log"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/events"
	l "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type Payload struct {
	Name string `json:"name"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.QueryStringParameters["name"]
	log.Printf("Name: %s", name)

	if 0 == len(name) {
		name = "None Set"
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(sess, &aws.Config{Region: aws.String("eu-west-2")})

	payload, err := json.Marshal(&Payload{ Name: name })
    if err != nil {
        fmt.Println("Error marshalling Payload")
        os.Exit(0)
    }

    result, err := client.Invoke(&lambda.InvokeInput{
    	FunctionName: aws.String("HttpInternalApiGo-dev-internalApi"),
    	Payload: payload,
    })
	if err != nil {
		fmt.Println("HttpInternalApiGo-dev-internalApi")
		os.Exit(0)
	}

	log.Print(result)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: string(result.Payload),
	}, nil
}

func main() {
	l.Start(Handler)
}