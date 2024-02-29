package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	target int
	output []int
}

func Test_twoSumBest(t *testing.T) {
	testcases := []TestCase{
		{nums: []int{2, 7, 11, 15}, target: 9, output: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, output: []int{1, 2}},
		{nums: []int{3, 2, 4}, target: 8, output: []int{-1, -1}},
	}
	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("run iteration number %d", i+1), func(test *testing.T) {
			output := twoSumBest(testcase.nums, testcase.target)
			if !reflect.DeepEqual(output, testcase.output) {
				test.Errorf("expected %+v, got %+v", testcase.output, output)
			}
		})
	}
}
