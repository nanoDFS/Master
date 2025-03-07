package utils

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/nanoDFS/Master/utils/config"
)

func InitLog() {
	file, err := os.Create(config.LoadConfig().Log.Path)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(file)

	styles := log.DefaultStyles()
	log.SetFormatter(log.TextFormatter)
	log.SetStyles(styles)
	log.SetLevel(log.InfoLevel)
}
