package main

import "fmt"

func main() {
	array := []int{0, 1, 0, 3, 12}
	 moveZeroes(array)
	//moveZeroes2(array)
	fmt.Println(array)
}

//Two poiter left and right(take all non-zero in leftside(swap) and increase leftpointer  )
func moveZeroes(nums []int) {
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[r], nums[l] = nums[l], nums[r]
			l++
		}
	}
}
func moveZeroes1(nums []int) {
	l := 0

	for r := range nums {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l++
		}
	}
	for ; l < len(nums); l++ {
		nums[l] = 0
	}
	return
}
func moveZeroes2(nums []int) {
	total_zeros := 0
	for i := 0; i < len(nums)-total_zeros; i++ {
		if nums[i] == 0 {
			total_zeros += 1
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			i = i - 1
		}
	}

}
