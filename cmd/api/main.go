package main

import (
	"log"
	"github.com/spf13/viper"
	"github.com/sablev/go-clean-architecture-std/internal/config"
	"github.com/sablev/go-clean-architecture-std/internal/server"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
