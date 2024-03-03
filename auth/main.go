package main

import (
	"log"

	"auth/internal/app"
)

const (
	appName = "auth"
)

func main() {
	if err := app.Run(appName); err != nil {
		log.Fatal(err)
	}
}
