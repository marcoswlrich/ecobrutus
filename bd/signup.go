package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/marcoswlrich/ecobrutus/models"
	"github.com/marcoswlrich/ecobrutus/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comecando Registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	seventy := "INSERT INTO users (User_email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "', '" + tools.DateMySQL() + "' )"
	fmt.Println(seventy)

	_, err = Db.Exec(seventy)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Execucao com sucesso ")
	return nil
}
