package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var firstBolean bool
	secondBollean := true
	if firstBolean || secondBollean {
		fmt.Println(firstBolean)
	}

	// Numerics int8, int16, int32, int64, int == OS в битах
	// uint8, uint16, uint32, uint64, uint
	var num1 int = 123
	num2 := 231
	fmt.Println("Sum of num and num2", num1+num2)

	// Вывод типа %T
	fmt.Printf("Type %T size of %d bytes\n", num2, unsafe.Sizeof(num2)) // unsafe.SizeOf в байтах (байт = 8 бит)

	// разные инты
	var n1 int64 = 5
	var n2 int32 = 10
	fmt.Printf("Sum of 5 and 10 = %d\n", int32(n1)+n2)

	// Numeric float32, float64. Нет Double
	var firstfloat float32 = 1.2
	var secondFloat float32 = 5.5
	fmt.Printf("Sum of float = %.3f\n", firstfloat+secondFloat)

	// Numeric Complex
	// firstComplex := 4 + 5i
	// secondComplex := 12 + 32i
	// fmt.Printf(firstComplex * secondComplex)

	// String
	name := "Viktor"
	lastname := "F3ss"
	concat := name + " " + lastname

	fmt.Printf("Full name: %s, len of it: %d\n", concat, len(concat))  // len() количество элементов
	fmt.Println("Amount of chars", name, utf8.RuneCountInString(name)) //Rune - один utf -ый символ
	fmt.Print(strings.Contains(name, "ktor"))                          // Поиск подстроки

	// Const
	const pi float32 = 3.1415

	fmt.Println(pi)

}
