package utils

import (
	"strings"
)

func StringCompare(a string, b string) bool {

	a = strings.ReplaceAll(strings.ToLower(a), " ", "")
	b = strings.ReplaceAll(strings.ToLower(b), " ", "")
	return a == b
}
