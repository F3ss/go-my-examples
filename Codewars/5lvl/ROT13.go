package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Rot13("EBG13 rknzcyr."))
	fmt.Println(Rot13("@[`{"))
	fmt.Println(Rot13("123"))
	fmt.Println(Rot13("How can you tell an extrovert from an\nintrovert at NSA? Va gur ryringbef,\ngur rkgebireg ybbxf ng gur BGURE thl'f fubrf."))
	fmt.Println("Ubj pna lbh gryy na rkgebireg sebz na\\nvagebireg ng AFN? In the elevators,\\nthe extrovert looks at the OTHER guy's shoes.")
	str1 := Rot13("How can you tell an extrovert from an\nintrovert at NSA? Va gur ryringbef,\ngur rkgebireg ybbxf ng gur BGURE thl'f fubrf.")
	str2 := "Ubj pna lbh gryy na rkgebireg sebz na\\nvagebireg ng AFN? In the elevators,\\nthe extrovert looks at the OTHER guy's shoes."
	fmt.Println(str1 == str2)
}

func Rot13(msg string) string {
	result := ""

	for _, char := range msg {
		if string(char) == "\n" {
			result += strings.ReplaceAll(string(char), "\n", "\\n")
		} else if string(char) == "\t" {
			result += strings.ReplaceAll(string(char), "\t", "\\t")
		} else if string(char) == "\a" {
			result += strings.ReplaceAll(string(char), "\a", "\\a")
		} else {
			swapChar(char, &result)
		}
	}

	return result
}

func swapChar(char int32, result *string) {
	var r string

	if char > 64 && char < 91 {
		if char+13 < 91 {
			r = string(char + 13)
		} else {
			r = string(char + 13 - 91 + 65)
		}
		*result += r
	} else if char > 96 && char < 123 {
		if char+13 < 123 {
			r = string(char + 13)
		} else {
			r = string(char + 13 - 123 + 97)
		}
		*result += r
	} else {
		*result += string(char)
	}
}
