//Написать программу, которая конкурентно рассчитает значение квадратов
//Чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
// ВАРИАНТ 2
package main

import (
	"fmt"
	"math"
)

func main() {
	ch := make(chan float64)
	numbers := []float64{2, 4, 6, 8, 10}
	for _, i := range numbers {
		go mathPow(i, ch)
		// вывод значений из канала
		fmt.Println(<-ch)
	}
}

// mathPow записывает результат в канал
func mathPow(number float64, ch chan float64) {
	ch <- math.Pow(number, 2)
}

