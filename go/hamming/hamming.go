// This function provides the hamming distance for two equal length DNA segments.
package hamming

import "errors"

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("DNA strings are not equal lengths")
	}

	var hamming_distance = 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hamming_distance += 1
		}
	}
	return hamming_distance, nil
}
