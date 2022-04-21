package main

import "fmt"

type P struct {
	X int
	Y int
}

func (p P) Abs() int {
	return p.X + p.Y
}

func main() {
	p := P{1, 2}
	fmt.Println(p.Abs())

}
