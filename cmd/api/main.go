package main

import (
	"log"

	"github.com/sablev/go-clean-architecture-std/internal/config"
	"github.com/sablev/go-clean-architecture-std/internal/server"

	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.New()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
