package main

import "fmt"

func main() {
	arr := []int{1234, 5678, 9012}
	num := 14690

	func(arr []int, num int) {
		for indI, i := range arr {
			for indN := indI + 1; indN < len(arr); indN++ {
				if i+arr[indN] == num {
					fmt.Println(indI, indN)
				}
			}
		}
	}(arr, num)

}
