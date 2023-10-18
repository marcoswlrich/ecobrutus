package main

import (
	"context"
	"errors"
	"fmt"
	"os"

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

	if !ValidateParameters() {
		fmt.Println("Erro nos parametros, deve-se enviar 'SecretName'")
		err := errors.New("Erro nos parametros deve-se enviar SecretName")
		return event, err
	}
}

func ValidateParameters() bool {
	var bringsParameters bool
	_, bringsParameters = os.LookupEnv("SecretName")
	return bringsParameters
}
