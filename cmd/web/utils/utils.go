package utils

import (
	"os"
)

const DEFAULT_PORT = "8080"

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	return ":" + port
}
