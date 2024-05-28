package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppName      string `env:"APP_NAME" env-default:"YT Time Tracker"`
	IsDebug      bool   `env:"APP_DEBUG" env-default:"false"`
	IsProduction bool   `env:"APP_PRODUCTION" env-default:"true"`
	Token        string `env:"TOKEN" env-required:"true"`
	ApiUrl       string `env:"YT_URL" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("get congig")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			var helpTex = "The Chexov"
			help, _ := cleanenv.GetDescription(instance, &helpTex)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance

}
