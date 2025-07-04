package utils

import (
	"fmt"
	"os"
)

func GetEnv(key string) (string, error) {
	value, ok := os.LookupEnv(key)

	if ok {
		return value, nil
	}

	return "", fmt.Errorf("environment variable %s is not set", key)
}