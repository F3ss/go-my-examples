package main

import (
	"reflect"
	"testing"
)

type testStruct struct {
	nums   []int
	target int
	result []int
}

func Test_TwoSum_ByMap(t *testing.T) {
	testValues := []testStruct{
		{nums: []int{1, 7, 11, 2}, target: 9, result: []int{1, 3}},
		{nums: []int{3, 3}, target: 6, result: []int{0, 1}},
		{nums: []int{3}, target: 6, result: nil},
		{nums: []int{3, 2, 4, -32, 54, 12}, target: 22, result: []int{3, 4}},
	}

	for _, test := range testValues {
		if !reflect.DeepEqual(TwoSum_ByMap(test.nums, test.target), test.result) {
			t.Errorf(
				"Result should be %v but result is %v\n",
				test.result,
				TwoSum_ByMap(test.nums, test.target))
		}
	}
}

func Test_TwoSum_ByTwoLoop(t *testing.T) {
	testValues := []testStruct{
		{nums: []int{1, 7, 11, 2}, target: 9, result: []int{1, 3}},
		{nums: []int{3, 3}, target: 6, result: []int{0, 1}},
		{nums: []int{3}, target: 6, result: nil},
		{nums: []int{3, 2, 4, -32, 54, 12}, target: 22, result: []int{3, 4}},
	}

	for _, test := range testValues {
		if !reflect.DeepEqual(TwoSum_ByTwoLoop(test.nums, test.target), test.result) {
			t.Errorf(
				"Result should be %v but result is %v\n",
				test.result,
				TwoSum_ByTwoLoop(test.nums, test.target))
		}
	}
}

func Benchmark_TwoSum_ByMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum_ByMap([]int{1, 7, 11, 2}, 9)
	}
}

func Benchmark_TwoSum_ByTwoLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum_ByTwoLoop([]int{1, 7, 11, 2}, 9)
	}
}

// Benchmark_TwoSum_ByMap-8   	 	1492880	      		730.6 ns/op	    2728 B/op	        3 allocs/op
// Benchmark_TwoSum_ByTwoLoop-8   	50080128	        36.90 ns/op	    16 B/op	       		1 allocs/op
