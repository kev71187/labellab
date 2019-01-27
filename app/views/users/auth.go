package users

import (
	"app/models"
	// "app/store"
	"fmt"
)

func Auth(auth string) {
	currentUser := models.CurrentUserAuth(auth)
	fmt.Println("Auth saved. You are now logged in")

	models.SetCurrentUser(currentUser)
}
