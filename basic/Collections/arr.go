package main

import "fmt"

func main() {
	// Array
	var arr [5]int
	fmt.Println(arr)
	arr[0], arr[2], arr[4] = 5, 3, 2
	arr1 := [2]int{1, 2}
	arrWithValues := [...]int{3, 4}
	fmt.Printf("After elements init: %v\n%v\n%v\n", arr, arr1, arrWithValues)

	floatArr := [...]float64{1.32, 2.32, 1.54, 65.2, 7.2}
	words := [2][2]string{
		{"Bob", "Dob"},
		{"Pod", "zxc"},
	}
	fmt.Println(floatArr, "\n", words)

	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Printf("%s ", val2)
		}
		fmt.Println()
	}

}
