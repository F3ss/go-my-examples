package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	var result string

	for _, word := range strings.Fields(str) {
		if unicode.IsLetter(rune(word[0])) {
			result += string(unicode.ToUpper(rune(word[0])))
		}
	}

	fmt.Println(result)
}
