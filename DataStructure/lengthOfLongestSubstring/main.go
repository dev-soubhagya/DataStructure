package main

import (
	"fmt"
)

func main() {
	// fmt.Println(time.Now())
	len := lengthOfLongestSubstring("abbabcda")
	fmt.Println(len)
	// len1 := lengthOfLongestSubstring1("abbabcda")
	// log.Println(len1, time.Now())
	// len2 := lengthOfLongestSubstring1("abba")
	// fmt.Println(len2)
}

func lengthOfLongestSubstring(s string) int { //abbabcda
	charLastIndex := make(map[rune]int)

	longestSubstringLength := 0
	currentSubstringLength := 0
	start := 0

	for index, character := range s {
		//fmt.Println("character ", character, "  index ", index)
		lastIndex, ok := charLastIndex[character]
		fmt.Println("lastindex ", lastIndex, " ispresent ", ok)
		if !ok || lastIndex < index-currentSubstringLength {
			currentSubstringLength++
		} else {
			if currentSubstringLength > longestSubstringLength {
				longestSubstringLength = currentSubstringLength
			}
			start = lastIndex + 1
			currentSubstringLength = index - start + 1
		}
		charLastIndex[character] = index
	}
	if currentSubstringLength > longestSubstringLength {
		longestSubstringLength = currentSubstringLength
	}
	return longestSubstringLength
}

// func lengthOfLongestSubstring1(s string) int { //abba
// 	seen := make(map[byte]int)
// 	maxLength := 0
// 	left := 0

// 	for right := 0; right < len(s); right++ {
// 		if _, ok := seen[s[right]]; ok { //a->0,
// 			// Move the left window boundary to the next index of the repeated character
// 			left = max(left, seen[s[right]]+1)
// 		}
// 		// Update the last seen index of the current character
// 		seen[s[right]] = right
// 		// Update the maximum length seen so far
// 		maxLength = max(maxLength, right-left+1)
// 	}

// 	return maxLength
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
