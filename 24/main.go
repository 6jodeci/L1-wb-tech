// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package main

import (
	"fmt"
	"math"
)

// описываем точку типом данных
type Point struct {
	x, y float64
}

// Values определяет значения точек
func Values(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance рассчитывет расстояние между точками по разнице координат
// d = sqrt(x2-x1)^2+(y2-y1)^2
func Distance(dot1, dot2 *Point) float64 {
	// возведение в квадрат
	dx := math.Pow(dot2.x-dot2.x, 2)
	dy := math.Pow(dot2.y-dot1.y, 2)
	// вычисление под корнем
	return math.Sqrt(dx + dy)
}

func main() {
	dot1 := Values(5, 5)
	dot2 := Values(1, 1)
	fmt.Println("Distance between X and Y: ", Distance(dot1, dot2))
}
