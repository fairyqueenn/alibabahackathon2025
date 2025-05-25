package helper

import (
	"regexp"
	"strings"
)

func ConvertToInLineQuery(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "  ", " ")

	re := regexp.MustCompile(` +\r?\n +`)
	return re.ReplaceAllString(s, " ")
}

func IsErrNoRows(s string) bool {
	return strings.Contains(s, "no rows")
}
