package env

import (
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	if err := godotenv.Load("./env.local"); err != nil {
		log.Fatal("Error loading env.local ./env.local")
	}
}
