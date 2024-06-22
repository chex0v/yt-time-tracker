package config

import (
	"errors"
	"log"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Token  string `yaml:"TOKEN" env:"TOKEN" env-required:"true"`
	ApiUrl string `yaml:"YT_URL" env:"YT_URL" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		var pathRc string
		if _, err := os.Stat("config.yaml"); errors.Is(err, os.ErrNotExist) {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				panic("Home dir not exist!")
			}
			if _, err := os.Stat(homeDir + "/.yttrc/config.yaml"); errors.Is(err, os.ErrNotExist) {
				panic("Config file not found")
			} else {
				pathRc = homeDir + "/.yttrc/config.yaml"
			}
		} else {
			pathRc = "config.yaml"
		}

		if pathRc == "" {
			panic("Config file not found")
		}
		if err := cleanenv.ReadConfig(pathRc, instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance

}
