package main

import "fmt"

//Example:
//Input : "soubhagya ranjan panda"
//Output : "panda ranjan soubhagya"
func main() {
	input := "chita ranjana biswal"
	temp := ""
	output := ""
	for i := 0; i < len(input); i++ {
		if string(input[i]) != "" {
			temp = temp + string(input[i])
		}
		if string(input[i]) == " " {
			output = temp + " " + output
			temp = ""
		}
		if i == len(input)-1 {
			output = temp + " " + output
		}
	}
	fmt.Println(output)
}

// Reverse the words in the given string
//Example:
//Input : "soubhagya ranjan panda"
//Output : "panda ranjan soubhagya"

/*
	var output string
	lists := strings.Split(input, " ")
	for i := len(lists) - 1; i >= 0; i-- {
		output += lists[i] + " "
	}
	fmt.Println(output)

*/

/*
	input := "chita ranjana biswal"
	temp := ""
	output := ""
	for i := 0; i < len(input); i++ {
		if string(input[i]) != "" {
			temp = temp + string(input[i])
		}
		if string(input[i]) == " " {
			output = temp + " " + output
			temp = ""
		}
		if i == len(input)-1 {
			output = temp + " " + output
		}
	}
	fmt.Println(output)
*/
