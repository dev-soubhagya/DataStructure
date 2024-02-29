package main

import "fmt"

func main() {
	Input := "Hello! How are you?" //you? are How Hello!
	fmt.Println(ReverseSentence(Input))
}

// ReverseSentence reverse a sentence
func ReverseSentence(sentence string) string {
	runes := []rune(sentence)
	size := len(runes)

	// Reverse entire string
	Reverse(runes, 0, size-1)
	fmt.Println("rev ", string(runes))
	start := 0
	end := 0

	for start < size && end < size {
		if runes[start] == ' ' {
			start++
			end++
			continue
		}

		if runes[end] != ' ' {
			end++
			continue
		}

		Reverse(runes, start, end-1)
		start = end
	}

	Reverse(runes, start, end-1)
	return string(runes)
}

// Reverse reverse a string
func Reverse(runes []rune, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

// Description
// Given a sentence, reverse it word by word.

// Example 1:
// Input:  "Hello! How are you?"
// Output: "you? are How Hello!"
