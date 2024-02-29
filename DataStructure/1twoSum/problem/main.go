package main

import (
	"fmt"
)

func main() {
	// array := []int{2, 7, 11, 19}
	// target := 9
	fmt.Println("Take the input array length ")
	var n int
	fmt.Scan(&n)
	array := make([]int, n)
	var target int
	fmt.Println("Type Input array ")
	for i := 0; i < len(array); i++ {
		_, err := fmt.Scan(&array[i])
		if err != nil {
			fmt.Println("err  in input array ", err)
		}
	}
	fmt.Println("arry is : ", array)
	fmt.Println("Target Value ")
	_, err1 := fmt.Scan(&target)
	if err1 != nil {
		fmt.Println("error in target scan ", err1)
	}
	//output_Indicies := twoSum(array, target)
	output_Indicies := twoSumBest(array, target)
	fmt.Println("OutPut : ", output_Indicies)
}

//O(n^2)
func twoSum(array []int, target int) []int {
	for first := 0; first < len(array); first++ {
		for second := first + 1; second < len(array); second++ {
			if first+second == target {
				return []int{first, second}
			}
		}
	}
	return []int{-1, -1}
}

//O(n)-HashMap Algorithm
func twoSumBest(array []int, target int) []int {
	// dataMap := map[int]int{}
	var dataMap map[int]int = make(map[int]int)
	for i, num := range array {
		diff := target - num
		if data, ok := dataMap[diff]; ok {
			fmt.Println("Found ", i, data)
			return []int{data, i}
		}
		dataMap[num] = i
	}
	return []int{-1, -1}
}
