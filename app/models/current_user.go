package models

import (
	"bytes"
	"encoding/json"
	"log"
)

type (
	CurrentUser struct {
		Id        uint64 `json:"id"`
		CreatedAt string `json:"created_at"`
		Email     string `json:"email"`
		AuthToken string `json:"auth_token"`
	}
	CurrentUserData struct {
		Data CurrentUser `json:"data"`
	}
	LoginUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	LoginData struct {
		Data LoginUser `json:"data"`
	}
	ErrorMessage struct {
		Message string `json:"message"`
	}
)

func CurrentUserAuth(auth string) CurrentUser {
	resp := RequestAuth(auth)
	var d CurrentUserData
	err := json.Unmarshal(resp, &d)
	if err != nil {
		log.Fatal(err)
	}
	writeUser(d.Data)
	return d.Data
}

func CurrentUserLogin(username string, password string) CurrentUser {
	reqBody := LoginData{Data: LoginUser{Username: username, Password: password}}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(reqBody)

	resp := Request("POST", "users/login", b)
	var d CurrentUserData
	err := json.Unmarshal(resp, &d)
	if err != nil {
		log.Fatal(err)
	}
	writeUser(d.Data)
	return d.Data
}
