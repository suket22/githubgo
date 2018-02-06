package main

import (
	"io/ioutil"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-lambda-go/events"
)

// Handler takes in an APIGatewayProxyRequest, processes it,
// and returns an APIGatewayProxyResponse.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	index, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(index),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}
