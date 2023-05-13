package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Could not load environment file. Error: %v", err)
	}

	var (
		env_value string
	)

	if value_retrieved, present := os.LookupEnv(key); !present {
		log.Fatal("Env value is empty")
	} else {
		env_value = value_retrieved
	}

	return env_value
}
