package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(line string) bool {
	uniques := make(map[rune]bool)
	for _, letter := range strings.ToLower(line) {
		if letter >= 'a' && letter <= 'z' {
				uniques[letter] = true
		}
	}
	return len(uniques) == 26
}
