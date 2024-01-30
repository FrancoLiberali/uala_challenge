package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/FrancoLiberali/uala_challenge/aws_lambda/handlers"
)

func main() {
	lambda.Start(handlers.HandleTweet)
}
