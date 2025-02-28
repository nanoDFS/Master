package secrets

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load("/Users/nagarajpoojari/Desktop/learn/nanoDFS/Master/.env")
	if err != nil {
		log.Fatalf("Error loading .env file, %v", err)
	}
}

func Get(key string) string {
	Load()
	return os.Getenv(key)
}
