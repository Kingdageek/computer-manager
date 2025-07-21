package helpers

import (
	"os"
	"strconv"
)

func GetEnv[T any](key string, defaultValue T) T {
	if value, exists := os.LookupEnv(key); exists {
		var result T
		switch any(defaultValue).(type) {
		case string:
			result = any(value).(T)
		case int:
			if v, err := strconv.Atoi(value); err == nil {
				result = any(v).(T)
			} else {
				result = defaultValue
			}
		case bool:
			if v, err := strconv.ParseBool(value); err == nil {
				result = any(v).(T)
			} else {
				result = defaultValue
			}
		default:
			result = defaultValue
		}
		return result
	}
	return defaultValue
}
