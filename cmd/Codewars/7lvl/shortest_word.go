package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps"))
}

func FindShort(s string) int {
	str := strings.Split(s, " ")
	min := len(str[0])

	for _, i := range str {
		if min > len(i) {
			min = len(i)
		}
	}
	return min
}
