package main

import "testing"

func TestParseNumber(t *testing.T) {
	var tests = []struct {
		Input    string
		Expected string
	}{
		{"p123456789", "123456789"},
		{"p123p123p", "123123"},
		{"abc123", "123"},
		{"ss 789p 88", "78988"},
	}

	for _, test := range tests {
		if output := parseNumber(test.Input); output != test.Expected {
			t.Errorf("expected was %v but actual shows %v", test.Expected, output)
		}
	}
}
