package main

import "fmt"

func main() {
	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Printf("The number %d is even\n", value)
	} else {
		fmt.Printf("The number %d is odd\n", value)
	}

}
