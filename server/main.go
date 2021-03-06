package main

import (
	"log"

	"innovateafitc/config"
	"innovateafitc/server"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatalln("Failed to parse config", err)
	}

	srv, err := server.New(cfg)

	if err != nil {
		log.Fatalln("Failed to create server", err)
	}

	srv.Start()
}
