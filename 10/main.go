// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
package main

import "fmt"

func main() {
	// array values
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// map store [] floats
	floatgroup := make(map[int][]float64)

	for _, value := range array {
		j := int(value/10) * 10
		floatgroup[j] = append(floatgroup[j], value)
	}
	fmt.Println(floatgroup)
}
