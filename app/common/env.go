package common

import (
	"errors"
	"os"
)

var availableEnvs map[string]bool

func init() {
	availableEnvs = map[string]bool{
		"dev":    true,
		"stage":  true,
		"prod":   false,
		"docker": true,
	}
}

func GetEnv() (*string, error) {
	env := os.Getenv("APP_MODE")
	if env == "" {
		env = "dev"
	}

	if value, ok := availableEnvs[env]; !ok || value == false {
		return nil, errors.New("env does not exist or not activated")
	}

	return &env, nil
}
