package main

import "fmt"

//mom
func main() {
	s := "21022012"
	// fmt.Println(CheckPallindrom(s))
	// fmt.Println(CheckPallindrom1(s))
	fmt.Println(CheckPallindrom2(s))
}
func CheckPallindrom(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func CheckPallindrom1(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func CheckPallindrom2(s string) bool {
	for i := 0; i < (len(s))/2; i++ {
		if s[i] != s[(len(s)-1)-i] {
			return false
		}
	}
	return true
}
