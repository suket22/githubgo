package main

import (
	lambdahandler "./lambdahandler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambdahandler.Handler)
}
