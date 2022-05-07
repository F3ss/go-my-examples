package main

import (
	"fmt"
)

func main() {
	fmt.Println(Two_Sum([]int{1, 7, 11, 2}, 9))
}

func Two_Sum(nums []int, target int) []int {
	mapa := make(map[int]int, 100)

	for index, i := range nums {
		num := target - i

		if _, has := mapa[num]; has == false {
			mapa[num] = index
			continue
		}
		return []int{index, mapa[num]}
	}
	fmt.Println(mapa)
	return nil
}

// func Two_Sum(nums []int, target int) []int {

// 	for i := 0; i < len(nums); i++ {
// 		for y := i + 1; y < len(nums); y++ {
// 			if nums[y] == target-nums[i] {
// 				return []int{i, y}
// 			}
// 		}
// 	}
// 	return nil
// }
