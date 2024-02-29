package main

import "fmt"

// Input:  "hello"
// Output: "olleh"

//Example:
//Input : "soubhagya ranjan panda"
//Output : "adnap najnar aygahbuos"
func main() {
	input := "soubhagya ranjan panda"
	output := ReverseString(input)
	fmt.Println(output)
}

//ReverseString reverse a text
func ReverseString(text string) string {
	runes := []rune(text)
	// fmt.Println(string(runes))
	Reverse(runes, 0, len(runes)-1)
	return string(runes)
}

//Reverse reverse runes
func Reverse(runes []rune, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}
