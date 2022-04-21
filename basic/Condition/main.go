package main

import (
	"fmt"
	"strings"
)

func strConcat() func(s string) string {
	str := ""
	return func(s string) string {
		str += s
		return str
	}
}

func main() {
	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Printf("The number %d is even\n", value)
	} else {
		fmt.Printf("The number %d is odd\n", value)
	}

	str := strConcat()
	var strForInput string
	for i := 0; i < 3; i++ {
		fmt.Scan(&strForInput)
		str(strForInput + "\n")
	}

	strForInput = str("")
	arrForCheck := strings.Split(strForInput, "\n")
	fmt.Println(arrForCheck)

	switch arrForCheck[2] {
	case "hello":
		fmt.Println("a")
	case "world":
		fmt.Println("b")
		fallthrough // Провалиться в оставшиеся кейсы в низ без проверки
	default:
		fmt.Println("Hm...")
	}
}
