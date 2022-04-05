package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var number int
	counter := make(map[int]int)
	fmt.Scanf("%d", &number)

	for _, num := range strings.Split(strconv.Itoa(number), "") {
		n, _ := strconv.Atoi(num)
		counter[n] = counter[n] + 1
	}
	fmt.Println(counter)
	printCounter(counter)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// printCounter печатает карту в формате
// key1:val1 key2:val2 ... keyN:valN
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
