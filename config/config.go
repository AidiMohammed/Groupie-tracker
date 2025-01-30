package config

import (
	"groupie-tracker/tools"
	"os"
)

type Config struct {
	APP_PORT string
	API_URL string
	BASE_URL string
}

func LoadConfig() (Config,error) {

	var myConfig Config

	err := tools.LoadEnvFile("../.env")
	if err != nil {
		return myConfig, err
	}

	myConfig.APP_PORT = os.Getenv("APP_PORT")
	myConfig.API_URL = os.Getenv("API_URL")
	myConfig.BASE_URL = os.Getenv("BASE_URL")

	return myConfig,nil
}