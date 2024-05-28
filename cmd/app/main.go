package main

import (
	"log"

	"github.com/chex0v/yt-time-tracker/internal/config"
)

func main() {

	log.Print("Load config")

	cfg := config.GetConfig()

	log.Print(cfg)

}
