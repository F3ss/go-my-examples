package main

import (
	"fmt"
	"log"
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
	if _, err := fmt.Scan(&value); err != nil {
		log.Fatalln(err)
	}

	if value%2 == 0 {
		fmt.Printf("The number %d is even\n", value)
	} else {
		fmt.Printf("The number %d is odd\n", value)
	}

	str := strConcat()
	var strForInput string
	for i := 0; i < 3; i++ {
		if _, err := fmt.Scan(&strForInput); err != nil {
			log.Fatalln(err)
		}
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
