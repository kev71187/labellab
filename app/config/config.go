package config

type (
	ApiConfig struct {
		Url         string
		AccessToken string
	}
)

// var BaseUrl = "http://nginx/api/"

// var BaseUrl = "http://localhost:3334/api/"

// var BaseUrl = "http://localhost:3333/api/"

var BaseUrl = "https://www.labellab.io/api/"
var AppConfig = ApiConfig{BaseUrl, ""}
