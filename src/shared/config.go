package shared

import (
	"fmt"
	"os"
	"strconv"
)

func MustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("%s env var is required, but wasn't set", key))
	}
	return value
}

func MustGetenvAsInt(key string) int {
	value := MustGetenv(key)
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("%s env var was %s, but int is required", key, value))
	}
	return intValue
}
