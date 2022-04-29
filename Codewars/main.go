package main

import (
	"fmt"
	"strings"
)

func ToCamelCase(s string) string {
	strings.ReplaceAll(s, "_", " ")

	return ""
}

func main() {
	fmt.Println()
}
