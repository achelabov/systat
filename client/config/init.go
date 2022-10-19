package server

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load("local.env"); err != nil {
		log.Fatal("no .env file found")
	}
}
