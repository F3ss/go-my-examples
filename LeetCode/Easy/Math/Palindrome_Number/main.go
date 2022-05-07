package main

import "fmt"

func main() {
	// fmt.Println(PalidromNumber(3))
	fmt.Println(PalidromNumberV2(121))
}

func PalidromNumber(x int) bool {
	var divider int = 1

	if x < 0 {
		return false
	}

	for i := 1; i < x; i = i * 10 {
		divider = i
	}

	if x%divider == 0 && x > 9 {
		divider = divider * 10
	}

	for i, y := divider, 10; i > 0; i, y = i/10, y*10 {
		if x%(i*10)/i != x%y/(y/10) {
			return false
		}
	}

	return true
}
