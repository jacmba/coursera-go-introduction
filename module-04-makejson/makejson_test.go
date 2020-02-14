package main

import "testing"

func TestMakeJson(t *testing.T) {
	testData := map[string]string{
		"name":    "John",
		"address": "a street",
	}

	result := makejson(testData)
	expect := "{\"address\":\"a street\",\"name\":\"John\"}"

	if result != expect {
		t.Fatal(result, "does not match", expect)
	}
}
