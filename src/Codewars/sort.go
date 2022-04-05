package codewars

import (
	"sort"
)

func TwoOldestAges(ages []int) [2]int {
	sort.Sort(sort.Reverse(sort.IntSlice(ages)))
	return [2]int{ages[1], ages[0]}
}
