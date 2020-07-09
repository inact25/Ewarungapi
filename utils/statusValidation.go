package utils

import (
	"strings"
)

func IsStatusValid(status string) bool {
	validation := strings.ToLower(status)
	if validation != "i" && validation != "a" {
		return false
	}
	return true
}
