package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		path        string
		search      string
		oldString   string
		newStringth string
	)

	fmt.Println("Где ищем исходя из дирректории с файлом?")
	fmt.Scan(&path)
	fmt.Println("Какие ищем файлы")
	fmt.Scan(&search)
	fmt.Println("Какую строку?")
	fmt.Scan(&oldString)
	fmt.Println("На что меняем??")
	fmt.Scan(&newStringth)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if strings.Contains(info.Name(), search) {

			read, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			newContents := strings.Replace(string(read), oldString, newStringth, -1)

			err = ioutil.WriteFile(path, []byte(newContents), 0)
			if err != nil {
				panic(err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Gotovo)")
}
