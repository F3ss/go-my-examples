package main

import (
	"fmt"
	"strings"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	var password string

	for {
		fmt.Println("Insert password: ")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Password is to easy. Try again.")
			continue
		} else {
			fmt.Printf("Nice password, really secret: %s\n", password)
			break
		}
	}

}
