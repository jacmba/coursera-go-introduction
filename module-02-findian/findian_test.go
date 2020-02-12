package main

import (
	"fmt"
	"testing"
)

type testData struct {
	input  string
	expect bool
}

func TestFindian(t *testing.T) {
	data := []testData{
		{"ian", true},
		{"Ian", true},
		{"iuiygaygn", true},
		{"I d skd a efju N", true},
		{"ihhhhhn", false},
		{"ina", false},
		{"xian", false},
	}

	for _, d := range data {
		name := fmt.Sprintf("%s should be %t", d.input, d.expect)
		t.Run(name, func(t *testing.T) {
			result := find(d.input)
			if result != d.expect {
				t.Fatal(result, "does not match", d.expect)
			}
		})
	}
}
