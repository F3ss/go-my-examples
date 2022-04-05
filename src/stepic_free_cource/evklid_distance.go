package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		x1 int
		y1 int
		x2 int
		y2 int
	)
	// считываем числа из os.Stdin
	// гарантируется, что значения корректные
	// не меняйте этот блок
	fmt.Scan(&x1, &y1, &x2, &y2)

	// рассчитайте d по формуле эвклидова расстояния
	// используйте math.Pow(x, 2) для возведения в квардрат
	// используйте math.Sqrt(x) для извлечения корня
	d := math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))

	fmt.Println(d)
}
