package models

import (
	"app/config"
	// "fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func decorate(req http.Request) {
	req.Header.Set("Authorization", Store.CurrentUser.AuthToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("UserAgent", "cli")
}

func IsOurUrl(url string) bool {
	switch url {
	case
		"http://nginx/api/",
		"http://localhost:3333/api/",
		"https://labellab.io/api":
		return true
	}
	return false
}

func successCode(code int) bool {
	switch code {
	case
		200,
		201,
		202,
		203,
		204:
		return true
	}
	return false
}

func RequestWithRetries(method string, url string, reader io.Reader, retry int) []byte {
	var u = config.AppConfig.Url + url
	req, _ := http.NewRequest(method, u, reader)
	decorate(*req)
	// fmt.Println(u)

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, _ := netClient.Do(req)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)

	if !successCode(resp.StatusCode) {
		if retry == 0 {
			log.Fatal(method + " Url " + u + " failed with: " + strconv.Itoa(resp.StatusCode))
		} else {
			return RequestWithRetries(method, url, reader, retry-1)
		}
	}

	if err != nil {
		if retry == 0 {
			log.Fatal(err)
		} else {
			return RequestWithRetries(method, url, reader, retry-1)
		}
	}

	return bytes
}

func Request(method string, url string, reader io.Reader) []byte {
	var u = config.AppConfig.Url + url
	req, _ := http.NewRequest(method, u, reader)
	decorate(*req)
	// fmt.Println(u)

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Do(req)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)

	if !successCode(resp.StatusCode) {
		log.Fatal(method + " Url " + u + " failed with: " + strconv.Itoa(resp.StatusCode))
	}

	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()
	return bytes
}

func RequestAuth(auth string) []byte {
	var u = config.AppConfig.Url + "current_user"
	req, _ := http.NewRequest("GET", u, nil)
	decorate(*req)
	req.Header.Set("Authorization", auth)

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Do(req)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)

	if !successCode(resp.StatusCode) {
		log.Fatal(" Url " + u + " failed with: " + strconv.Itoa(resp.StatusCode))
	}

	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()
	return bytes
}
