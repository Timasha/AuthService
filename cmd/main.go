package main

import (
	"log"
	"os"

	"github.com/Timasha/AuthService/internal/app"
	executor "github.com/Timasha/AuthService/utils/app"
	"github.com/Timasha/AuthService/utils/config"
)

func main() {
	key, ok := os.LookupEnv("CONFIG")
	if !ok {
		log.Fatal("Cant read config path")
	}

	cfg, err := config.ReadConfigJSON(key)
	if err != nil {
		log.Fatal(err)
	}

	application := app.New(*cfg)

	err = executor.Run(application)
	if err != nil {
		log.Fatal(err)
	}
}
