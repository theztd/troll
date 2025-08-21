package libs

import (
	"os"
	"strconv"
)

// getnev like in Python
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func GetEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	num, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return num
}
