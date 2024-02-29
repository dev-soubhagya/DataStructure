package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{-4, -1, 0, 3, 10}
	SortedSquaredArray(array)
	fmt.Println("squares of a sorted array is : ", array)
}
func SortedSquaredArray2(nums []int) []int {
	res := []int{}
	l, r := 0, len(nums)-1
	for l = 0; l <= r; l++ {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			res = append(res, nums[l]*nums[l])
			l += 1
		} else {
			res = append(res, nums[r]*nums[r])
			r -= 1
		}
	}
	return res
}

//SortedSquaredArray takes an array of integers and returns the sorted squared result
func SortedSquaredArray1(array []int) []int {
	result := make([]int, len(array))
	for i := range array {
		result[i] = array[i] * array[i]
	}
	sort.Ints(result)
	return result
}

//SortedSquaredArray takes an array of integers and returns the sorted squared result
func SortedSquaredArray(array []int) []int {
	result := make([]int, len(array))
	leftValueIdx := 0
	rightValueIdx := len(array) - 1

	for currIdx := len(array) - 1; currIdx >= 0; currIdx-- {
		leftValue := array[leftValueIdx]
		rightValue := array[rightValueIdx]

		if abs(rightValue) > abs(leftValue) {
			result[currIdx] = rightValue * rightValue
			rightValueIdx -= 1
		} else {
			result[currIdx] = leftValue * leftValue
			leftValueIdx += 1
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
