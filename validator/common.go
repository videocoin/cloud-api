package validator

import "strings"

func extractValueFromTag(tag string) string {
	values := strings.Split(tag, ",")
	return values[0]
}
