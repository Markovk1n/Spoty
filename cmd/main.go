package main

import (
	"log"

	"github.com/markovk1n/spoty/configs"
	"github.com/markovk1n/spoty/internal/app"
)

func main() {
	cfg, err := configs.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(&cfg)
}
