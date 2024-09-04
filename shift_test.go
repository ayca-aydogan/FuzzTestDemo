package main

import (
	"testing"
	"unicode/utf8"
)

var testcases = []struct {
	input    string
	shift    int
	expected string
}{
	{
		input:    "",
		shift:    7,
		expected: "",
	},
	{
		input:    "$%~ghff",
		shift:    0,
		expected: "$%~ghff",
	},
	{
		input:    "exa4mple!!!!",
		shift:    7,
		expected: "e!!!!exa4mpl",
	},
	{
		input:    "Hello, world!!",
		shift:    15,
		expected: "ello, world!!H",
	},
	{
		input:    "00",
		shift:    -95,
		expected: "00",
	},
	{
		input:    "0Ÿ",
		shift:    7,
		expected: "Ÿ0",
	},
	{
		input:    "0Ÿ",
		shift:    2,
		expected: "0Ÿ",
	},
}

func TestShiftLeft(t *testing.T) {
	for _, tc := range testcases {
		output := ShiftLeft(tc.input, tc.shift)
		if tc.expected != output {
			t.Fatalf("expected: %v, got: %v", tc.expected, output)
		}
	}
}

func FuzzShiftLeft(f *testing.F) {
	for _, tc := range testcases {
		f.Add(tc.input, tc.shift)
	}
	f.Fuzz(func(t *testing.T, input string, shift int) {
		if !utf8.Valid([]byte(input)) {
			t.Skip()
		}
		output := ShiftLeft(input, shift)

		// output da geçerli utf8 olmalı
		if !utf8.Valid([]byte(output)) {
			t.Skip()
		}

		// Output ve inputun karakter sayıları aynı olmalı
		lenInput := utf8.RuneCountInString(input)
		lenOutput := utf8.RuneCountInString(output)
		if lenOutput != lenInput {
			t.Fatalf("num of char in input: %d\ninput: %s\n"+
				"num of char in output: %d"+
				"output: %s\nshift: %d\n",
				lenInput, input, lenOutput, output, shift)
		}

		// Inputu kendi uzunluğu kadar kaydırdığımda kendisini elde etmeliyim
		output2 := ShiftLeft(input, lenInput)
		if input != output2 {
			t.Fatalf("expected: %s, got: %s", input, output2)
		}
	})
}
