package appenv

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lookupEnvOrPanic(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Errorf("please set environment variable %s", key))
	}
	return value
}

func loadBool(key string) bool {
	value := lookupEnvOrPanic(key)
	ret, err := strconv.ParseBool(value)
	if err != nil {
		panic(fmt.Errorf("failed to parse environment variable (%s) due to %w", key, err))
	}
	return ret
}

func loadInt(key string) int {
	value := lookupEnvOrPanic(key)
	ret, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Errorf("failed to parse environment variable (%s) due to %w", key, err))
	}
	return ret
}

func loadString(key string) string {
	return lookupEnvOrPanic(key)
}

func loadStringSlice(key string) []string {
	value := lookupEnvOrPanic(key)
	return strings.Split(value, ",")
}
