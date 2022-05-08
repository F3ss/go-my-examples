package main

import (
	"fmt"
)

func main() {
	fmt.Println(TwoSum_ByMap([]int{1, 7, 11, 2}, 9))
	fmt.Println(TwoSum_ByMap([]int{3, 3}, 6))
}

func TwoSum_ByMap(nums []int, target int) []int {
	var num int
	mapa := make(map[int]int, 100)

	for i := 0; i < len(nums); i++ {
		num = target - nums[i]

		if _, has := mapa[num]; !has {
			mapa[num] = i
		}

		if _, has := mapa[nums[i]]; !has || has && mapa[nums[i]] == i {
			continue
		}
		return []int{mapa[target-num], i}
	}

	return nil
}

func TwoSum_ByTwoLoop(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for y := i + 1; y < len(nums); y++ {
			if nums[y] == target-nums[i] {
				return []int{i, y}
			}
		}
	}
	return nil
}
