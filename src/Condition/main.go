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

	var value1 int
	fmt.Scan(&value1)

	switch value1 {
	case 10:
		fmt.Println("a")
	case 20:
		fmt.Println("b")
		// fallthrough выполняет все кейсы включая дефолтный, которые идут ниже не проверяя значения
	default:
		fmt.Println("c")
	}
}
