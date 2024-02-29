package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	merge1(a, 3, b, 3)
	// merge(a, 3, b, 3)
	fmt.Println(a)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for k := m + n - 1; k >= 0; k-- {
		if i >= 0 && j >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else if j >= 0 {
			nums1[k] = nums2[j]
			j--
		}
	}
}
func merge1(nums1 []int, m int, nums2 []int, n int) {
	for t := m + n - 1; t >= 0; t-- {
		if m == 0 || n == 0 {
			if m == 0 {
				nums1[t] = nums2[n-1]
				n--
			} else {
				nums1[t] = nums1[m-1]
				m--
			}
			continue
		}
		if nums1[m-1] > nums2[n-1] {
			nums1[t] = nums1[m-1]
			m--
		} else {
			nums1[t] = nums2[n-1]
			n--
		}
	}
}
