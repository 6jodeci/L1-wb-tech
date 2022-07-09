// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
// ВАРИАНТ 1
package main

import (
	"fmt"
	"math"
)

func main() {
	ch := make(chan float64)
	numbers := []float64{2, 4, 6, 8, 10}
	var sum float64

	for _, i := range numbers {
		go mathPow(ch, i)
		// Запись суммы в канал
		sum += <- ch
	}
	// Вывод суммы квадратов
	fmt.Println(sum)
}

// mathPow записывает результат в канал
func mathPow(ch chan float64, numbers float64) {
	ch <- math.Pow(numbers, 2)
}
