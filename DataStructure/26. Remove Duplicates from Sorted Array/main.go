package main

import "fmt"

func main() {
	array := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// array := []int{1, 1, 2}
	fmt.Println(removeDuplicates(array))
}
func removeDuplicates(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return ln
	}

	j := 0 // points to  the index of last filled position
	for i := 1; i < ln; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	fmt.Println(nums)
	return j + 1
}
