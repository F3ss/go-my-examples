package main

import (
	"fmt"
)

func main() {
	fmt.Println(DecodeRoman("MMMM"))
}

func DecodeRoman(romanNum string) int {

	beforeTen := []string{
		"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
		"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC",
		"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM",
		"", "M", "MM", "MMM", "MMMM",
	}

	result := ""

	for i := 0; i < 5000; i++ {
		result = ""
		switch {
		case i >= 1000:
			result += beforeTen[i/1000+30]
			fallthrough
		case i >= 100:
			result += beforeTen[i%1000/100+20]
			fallthrough
		case i >= 10:
			result += beforeTen[i%100/10+10]
			fallthrough
		case i >= 1:
			result += beforeTen[i%10]
		}

		if romanNum == result {
			return i
		}
	}

	return 0
}
