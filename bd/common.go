package bd

import (
	"os"

	"github.com/marcoswlrich/ecobrutus/models"
	"github.com/marcoswlrich/ecobrutus/secretm"
)

var (
	SecretModel models.SecretRDSJson
	err         error
)

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
