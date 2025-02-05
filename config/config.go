package config

import (
	"groupie-tracker/tools"
	"os"
)

type Config struct {
	AppPort string
	ApiUrl string
	BaseUrl string
}

func LoadConfig() (Config,error) {

	var myConfig Config

	err := tools.LoadEnvFile("../.env")
	if err != nil {
		return myConfig, err
	}

	myConfig.AppPort = os.Getenv("APP_PORT")
	myConfig.ApiUrl = os.Getenv("API_URL")
	myConfig.BaseUrl = os.Getenv("BASE_URL")

	return myConfig,nil
}