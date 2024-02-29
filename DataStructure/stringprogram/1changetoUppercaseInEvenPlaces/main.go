package main

import (
	"fmt"
	"strings"
)

//Example:
//Input : "soubhagya"
//Output : "sOuBhAgYa"
func main() {
	data := "Hiii"
	var temp string
	for i := 0; i < len(data); i++ {
		if (i+1)%2 == 0 {
			temp = temp + strings.ToUpper(string(data[i]))
		} else {
			temp = temp + strings.ToLower(string(data[i]))
		}
	}
	fmt.Println(temp)
}

//change and print the given string to uppercase in even number of places and odd number of place to lower
//index start from 1
//Example:
//Input : "soubhagya"
//Output : "sOuBhAgYa"

/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "soubhagya"
	for i, s := range data {
		if i%2 != 0 {
			fmt.Printf("%s", strings.ToUpper(string(s)))
		} else {
			fmt.Printf("%s", strings.ToLower(string(s)))
		}
	}
}
*/

/*
var temp string
for i := 0; i < len(data); i++ {
	if (i+1)%2 == 0 {
		temp = temp + strings.ToUpper(string(data[i]))
	} else {
		temp = temp + string(data[i])
	}
}
fmt.Println(temp)

*/
