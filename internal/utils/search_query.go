package utils

import (
	"regexp"
	"strings"
)

func MakeSearchQuery(query string) string {
	re := regexp.MustCompile(`\w+`)
	words := re.FindAllString(query, -1)
	return strings.Join(words, " | ")
}
