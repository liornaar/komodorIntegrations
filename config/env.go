package config

import (
	"os"
	"strconv"
	"strings"
)

func String(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func StringSlice(key string, defaultValue []string) []string {
	if value, exists := os.LookupEnv(key); exists {
		return split(value)
	}
	return defaultValue
}

func Bool(key string, defaultValue bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseBool(value); err == nil {
			return result
		}
	}
	return defaultValue
}

func Int(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.Atoi(value); err == nil {
			return result
		}
	}
	return defaultValue
}

func Int32(key string, defaultValue int32) int32 {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseInt(value, 10, 32); err == nil {
			return int32(result)
		}
	}
	return defaultValue
}

func Int64(key string, defaultValue int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseInt(value, 10, 64); err == nil {
			return result
		}
	}
	return defaultValue
}

func Float32(key string, defaultValue float32) float32 {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseFloat(value, 32); err == nil {
			return float32(result)
		}
	}
	return defaultValue
}

func Float64(key string, defaultValue float64) float64 {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseFloat(value, 64); err == nil {
			return result
		}
	}
	return defaultValue
}

func split(value string) []string {
	var result []string
	values := strings.Split(value, ",")
	for _, v := range values {
		result = append(result, strings.TrimSpace(v))
	}
	return result
}
