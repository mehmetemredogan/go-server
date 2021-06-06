package boot

import (
	"github.com/joho/godotenv"
	"log"
)

func EnvLoader() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file error")
	}
}