package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvVar(key string) string{

	err:=godotenv.Load()

	if err!=nil {
		log.Fatal(envError)
	}

	return os.Getenv(key)
}