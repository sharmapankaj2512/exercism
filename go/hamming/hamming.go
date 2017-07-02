package hamming

import "errors"

const testVersion = 5

func Distance(a, b string) (int, error) {
	hammingDistance := 0

	if len(a) != len(b) {
		return 0, errors.New("")
	}

	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}

