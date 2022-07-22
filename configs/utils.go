package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(envError)
	}

	return os.Getenv(key)
}
