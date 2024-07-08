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
	Tasks  []Task `yaml:"TASKS"`
	Types  []Type `yaml:"TYPES"`
}

type Task struct {
	Key  string `yaml:"key"`
	Task string `yaml:"value"`
}

type Type struct {
	Key  string `yaml:"key"`
	Type string `yaml:"value"`
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

func (c *Config) TaskNumber(taskNumberFromConsole string) string {
	for _, t := range c.Tasks {
		if t.Key == taskNumberFromConsole {
			return t.Task
		}
	}
	return taskNumberFromConsole
}

func (c *Config) TypeId(codeType string) string {
	for _, t := range c.Types {
		if t.Key == codeType {
			return t.Type
		}
	}
	return ""
}
