package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// ------------------------------------------------- READ -------------------------------------------------
	content, err := ioutil.ReadFile("./example")

	if err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	fmt.Println(string(content))

	// ------------------------------------------------- READ -------------------------------------------------

	file, err := os.Open("./writed_file")

	if err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		fmt.Println(scan.Text())
	}

	// ------------------------------------------------- WRITE -------------------------------------------------

	f, err := os.Create("writed_file")

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	defer f.Close()

	_, err2 := f.WriteString("Hello World!\n")

	
	if err2 != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Println("Done")
}
