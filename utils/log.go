package utils

import (
	"github.com/charmbracelet/log"
)

func InitLog() {
	styles := log.DefaultStyles()
	log.SetStyles(styles)
	log.SetLevel(log.DebugLevel)
}
