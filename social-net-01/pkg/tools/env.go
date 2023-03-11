package tools

import "os"

func GetEnvWithDefault(name, defaultValue string) string {
	v, exists := os.LookupEnv(name)
	if !exists {
		return defaultValue
	}

	return v
}
