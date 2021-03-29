package utils

import (
	"os"
	"path"
	"strings"
)

const DEFAULT_PORT = "8080"

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	return ":" + port
}

func GetProjectRoot() (string, error) {
	workingDir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	if strings.Contains(workingDir, "cmd/web") {
		return path.Join(workingDir, "../../"), nil
	}

	return workingDir, nil
}
