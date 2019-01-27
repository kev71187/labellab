package users

import (
	"app/models"
	// "app/store"
	"fmt"
	"github.com/howeyc/gopass"
	"log"
)

func Login() {
	var username string
	fmt.Println("username/email:")
	fmt.Scanf("%v", &username)
	fmt.Println("password:")
	pass, err := gopass.GetPasswd()
	var password = string(pass)
	if err != nil {
		log.Fatal("Input authorization failed")
	}

	currentUser := models.CurrentUserLogin(username, password)
	models.SetCurrentUser(currentUser)
}
