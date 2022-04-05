package main

import (
	"fmt"
	"sort"
	"strings"
)

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

func main() {
	// m := make(map[string]int)
	// m["key"] = 7
	// m["other"] = 13

	// n := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println("map:", n)
	// // map: map[bar:2 foo:1]

	// _, ok := m["other"]
	// fmt.Println("has other:", ok)
	// // has other: false

	var number string
	fmt.Scanf("%s", &number)

	counter := make(map[int]int)

	fmt.Println(strings.Split(number, ""))
	for _, num := range strings.Split(number, "") {
		counter[int(num)] = counter[int(num)] + 1
	}

	printCounter(counter)

}
