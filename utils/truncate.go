package utils

import "strings"

func Truncate(str string, length int) string {
	trimmed := strings.TrimSpace(str)
	if len(trimmed) <= length {
		return trimmed
	}
	return strings.Join([]string{trimmed[:(length - 3)], "…"}, "")
}
