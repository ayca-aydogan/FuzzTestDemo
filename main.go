package main

import "fmt"

func main() {
	input := "this is the string to be shifted"
	shift := 4
	expectedOutput := " is the string to be shiftedthis"
	output := ShiftLeft(input, shift)

	fmt.Printf("input: %s\nshift: %d\nexpected output: %s\noutput: %s\nis output correct: %t",
		input,
		shift,
		expectedOutput,
		output,
		expectedOutput == output)
}
