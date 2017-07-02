package raindrops

import (
	"strconv"
	"bytes"
)

const testVersion = 3

func Convert(n int) string {
	var buffer bytes.Buffer

	if n%3 == 0 {
		buffer.WriteString("Pling")
	}
	if n%5 == 0 {
		buffer.WriteString("Plang")
	}
	if n%7 == 0 {
		buffer.WriteString("Plong")
	}
	if buffer.Len() == 0 {
		buffer.WriteString(strconv.Itoa(n))
	}

	return buffer.String()
}
