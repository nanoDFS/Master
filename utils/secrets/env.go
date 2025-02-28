package secrets

import (
	"os"
	"sync"

	"github.com/charmbracelet/log"

	"github.com/joho/godotenv"
)

var mu sync.RWMutex

func Load() {
	mu.Lock()
	err := godotenv.Load("/Users/nagarajpoojari/Desktop/learn/nanoDFS/Master/.env")
	if err != nil {
		log.Fatalf("Error loading .env file, %v", err)
	}
	mu.Unlock()
}

func Get(key string) string {
	Load()
	return os.Getenv(key)
}
