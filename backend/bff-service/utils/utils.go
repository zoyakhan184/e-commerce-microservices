package utils

import (
	"strconv"
)

// ParseInt converts a string to int and falls back to defaultVal if conversion fails.
func ParseInt(s string, defaultVal int) int {
	if val, err := strconv.Atoi(s); err == nil {
		return val
	}
	return defaultVal
}
