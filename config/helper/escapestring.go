package helper

import (
	"html"
	"regexp"
	"strings"
)

func EscapeStrings(kata string) string {
	return html.EscapeString(strings.TrimSpace(kata))
}

func NonLatinChar(kata string) []string {
	words := regexp.MustCompile("[\\p{L}\\d_]+")
	return words.FindAllString(kata, -1)
}
