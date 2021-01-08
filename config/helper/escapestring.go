package helper

import (
	"html"
	"strings"
)

func EscapeStrings(kata string) string {
	return html.EscapeString(strings.TrimSpace(kata))
}
