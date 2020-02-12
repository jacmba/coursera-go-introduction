package main

import (
	"fmt"
	"testing"
)

type testData struct {
	input  string
	expect int
}

func TestTruncFromString(t *testing.T) {
	data := []testData{
		{"10", 10},
		{"27.35", 27},
		{"30.1", 30},
		{"108.9", 108},
	}

	for _, d := range data {
		testName := fmt.Sprintf("'%s' should be %d", d.input, d.expect)
		t.Run(testName, func(t *testing.T) {
			result := truncate(d.input)
			if result != d.expect {
				t.Fatal(result, "does not match", d.expect)
			}
		})
	}
}
