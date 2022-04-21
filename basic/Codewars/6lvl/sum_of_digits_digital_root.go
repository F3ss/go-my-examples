package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 992
	str := strconv.Itoa(num)
	sum := 0

	for len(str) > 1 {
		for _, i := range str {
			num, _ = strconv.Atoi(string(i))
			sum += num
		}
		str = strconv.Itoa(sum)
		sum = 0
	}

	var result, _ = strconv.Atoi(str)
	fmt.Println(result)
}
