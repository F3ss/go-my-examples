package main

import "fmt"

func main() {
	// Слайс - набор ссылок на массив.
	// Если слайс делается с массива, при изменении слайса массив будет также изменён
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	slice1 := slice[0:2]
	fmt.Println(slice1)
	for index, value := range slice {
		fmt.Println(index, value)
	}

	// Capasity (cap) - это значение показывающее сколько элементов можно добавить в срез без выделения доп памяти в нижележащий массив.
	sliceOfString := []string{"One", "Two", "Three"}
	fmt.Println("sliceIfString:", sliceOfString, "len:", len(sliceOfString), cap(sliceOfString))
	sliceOfString = append(sliceOfString, "Four")
	fmt.Println("sliceIfString:", sliceOfString, "len:", len(sliceOfString), cap(sliceOfString))
}
