package utils

import (
	"os"
)

const ENV string = "ddd-env"

func GetEnvironment() string {
	defaultEnv := "dev"
	if env := os.Getenv(ENV); env != "" {
		return env
	}
	return defaultEnv
}
