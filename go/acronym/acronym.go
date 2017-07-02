package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Abbreviate(text string) string {
	words := regexp.MustCompile(`\s|-`).Split(text, -1)
	accronym := make([]byte, len(words))
	for i, word := range words {
		accronym[i] = word[0]
	}
	return strings.ToUpper(string(accronym))
}
