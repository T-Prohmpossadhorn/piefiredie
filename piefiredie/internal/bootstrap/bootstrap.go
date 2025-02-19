package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type EnvStruct struct {
	AppName string `mapstructure:"APP_NAME"`

	AppEnv        string `mapstructure:"APP_ENV"`
	GinMode       string `mapstructure:"GIN_MODE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`

	DataUrl string `mapstructure:"DATA_URL"`
}

var Env EnvStruct

func LoadEnv() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if Env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}
}
