package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}

func HighAndLow(in string) string {
	numbers := strings.Split(in, " ")
	var highAndLow []int

	for _, num := range numbers {
		n, _ := strconv.Atoi(num)
		highAndLow = append(highAndLow, n)
	}

	sort.Ints(highAndLow)
	return fmt.Sprintf("%d %d", highAndLow[len(highAndLow)-1], highAndLow[0])
}
