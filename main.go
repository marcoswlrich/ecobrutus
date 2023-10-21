package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	"github.com/marcoswlrich/ecobrutus/awsgo"
	"github.com/marcoswlrich/ecobrutus/bd"
	"github.com/marcoswlrich/ecobrutus/models"
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

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Erro ao ler Secret " + err.Error())
		return event, err
	}

	err = bd.SignUp(data)
	return event, err
}

func ValidateParameters() bool {
	var bringsParameters bool
	_, bringsParameters = os.LookupEnv("SecretName")
	return bringsParameters
}
