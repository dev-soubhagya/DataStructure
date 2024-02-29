package main

import (
	"fmt"
	"strings"
)

// aygahbuoS  najnaR adnaP
func main() {
	s := "Soubhagya Ranjan Panda"
	//Reverse(s)
	Reverse1(s)
}
func Reverse(s string) {
	var s1 string
	data := strings.Split(s, " ")
	for _, d := range data {
		for i := len(d) - 1; i >= 0; i-- {
			s1 += string(d[i])
		}
		s1 += " "
	}
	fmt.Println(s1)
}
func Reverse1(s string) {
	var temp string
	var output string
	for i := 0; i < len(s); i++ {
		if string(s[i]) != "" {
			temp = string(s[i]) + temp
		}
		if string(s[i]) == " " {
			output += temp + " "
			temp = ""
		}
		if i == len(s)-1 {
			output += temp
		}
	}
	fmt.Println(output)
}
