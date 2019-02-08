package config

type (
	ApiConfig struct {
		Url         string
		AccessToken string
	}
)

var Env string
var BaseUrl = "https://www.labellab.io/api/"

func load() ApiConfig {
	if Env == "local" {
		BaseUrl = "http://nginx/api/"
	}
	return ApiConfig{BaseUrl, ""}
}

var AppConfig = load()
