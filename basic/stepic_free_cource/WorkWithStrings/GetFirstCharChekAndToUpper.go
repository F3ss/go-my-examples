package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var abbr string

	rdr := bufio.NewReader(os.Stdin)
	phrase, _ := rdr.ReadString('\n')

	for _, word := range strings.Fields(phrase) {
		latter, _, _, err := strconv.UnquoteChar(word, 1)

		if err != nil {
			log.Fatalln(err)
		}

		if unicode.IsLetter(latter) {
			abbr += strings.ToUpper(string(latter))
		}

	}

	fmt.Println(abbr)
}
