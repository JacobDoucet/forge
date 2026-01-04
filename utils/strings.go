package utils

import (
	"regexp"
	"strings"
)

func SC(input string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	return strings.ToLower(re.ReplaceAllString(input, "${1}_${2}"))
}

func KC(input string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	return strings.ToLower(re.ReplaceAllString(input, "${1}-${2}"))
}

func UKC(input string) string {
	return strings.ToUpper(KC(input))
}

func LCC(input string) string {
	input = UCC(input)
	return strings.ToLower(input[:1]) + input[1:]
}

func UCC(input string) string {
	input = strings.ReplaceAll(input, "-", "_")
	parts := strings.Split(input, "_")
	for i := range parts {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}
