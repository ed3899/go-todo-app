package config

import (
	"log"
	"os"
)

func Config(key string) string {
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
