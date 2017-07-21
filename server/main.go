package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/pbtrung/Skypiea/server/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Init()
	StartServer()
}
