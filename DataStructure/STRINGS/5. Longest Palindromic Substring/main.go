package main

import "fmt"

func main() {
	// s := "babad"
	s := "abc"
	var s1 string
	i := 0
	j := len(s) - 1
	for i < j {
		if ispalllindrom(s[i:j+1]) && len(s1) > len(s) {
			s1 = s[i : j+1]
		} else if ispalllindrom(s[i : j+1]) {
			j--
		} else {
			i++
		}
	}
	fmt.Println("longeset ", s1)

}
func ispalllindrom(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
