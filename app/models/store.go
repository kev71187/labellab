package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io"
	"io/ioutil"
	// "log"
	// "os"
)

type GlobalStore struct {
	CurrentUser CurrentUser `json:"current_user"`
}

func SetCurrentUser(currentUser CurrentUser) {
	Store.CurrentUser = currentUser
	save()
}

func save() bool {
	configJson, _ := json.Marshal(Store)
	err := ioutil.WriteFile(StorePath, configJson, 0744)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func load() GlobalStore {
	raw, err := ioutil.ReadFile(StorePath)
	if err != nil {
		return GlobalStore{}
	}
	var c GlobalStore
	json.Unmarshal(raw, &c)
	return c
}

var dir, _ = homedir.Dir()
var StorePath = (dir + "/.labellab.json")
var Store = load()

func stuff() []byte {
	return []byte("c048e68f319b43d3e50262c9df07aa89")
}

func Encrypt(plaintext []byte) ([]byte, error) {
	key := stuff()
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	key := stuff()
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func ReadUser() {
}

func writeUser(user CurrentUser) {
	SetCurrentUser(user)
}
