package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	"github.com/marcoswlrich/ecobrutus/awsgo"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(
	ctx context.Context,
	event events.CognitoEventUserPoolsPostConfirmation,
) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitAWS()
}
