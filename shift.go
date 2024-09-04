package main

func ShiftLeft(input string, shift int) string {
	if shift < 0 {
		return input
	}
	inputBytes := []rune(input)
	lenInput := len(inputBytes)
	res := []rune{}
	for i := 0; i < lenInput; i++ {
		resIndex := (i + shift) % lenInput
		res = append(res, inputBytes[resIndex])
	}
	return string(res)
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

// first
// func ShiftLeft(input string, shift int) string {
// 	inputBytes := []byte(input)
// 	lenInput := len(inputBytes)
// 	res := []byte{}
// 	for i := 0; i < lenInput; i++ {
// 		resIndex := (i + shift) % lenInput
// 		res = append(res, inputBytes[resIndex])
// 	}
// 	return string(res)
// }

// second
// func ShiftLeft(input string, shift int) string {
// 	if shift < 0 {
// 		return input
// 	}
// 	inputBytes := []byte(input)
// 	lenInput := len(inputBytes)
// 	res := []byte{}
// 	for i := 0; i < lenInput; i++ {
// 		resIndex := (i + shift) % lenInput
// 		res = append(res, inputBytes[resIndex])
// 	}
// 	return string(res)
// }

// third
// func ShiftLeft(input string, shift int) string {
// 	if shift < 0 {
// 		return input
// 	}
// 	inputBytes := []rune(input)
// 	lenInput := len(inputBytes)
// 	res := []rune{}
// 	for i := 0; i < lenInput; i++ {
// 		resIndex := (i + shift) % lenInput
// 		res = append(res, inputBytes[resIndex])
// 	}
// 	return string(res)
// }
