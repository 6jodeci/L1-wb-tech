//Написать программу, которая конкурентно рассчитает значение квадратов 
//Чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
// ВАРИАНТ 1
package main

import (
	"fmt"
	"math"
	"sync"
)

var wg sync.WaitGroup

func main() {
	numbers := []float64{2, 4, 6, 8, 10}
	for _, i := range numbers {
		// Увеличение WaitGroup счетчика
		wg.Add(1)
		go mathPow(i)
	}
	// Ожидаем завершения
	wg.Wait()
}

// mathPow выводит число
func mathPow(number float64) {
	fmt.Println(math.Pow(number, 2))
	// Уменьшение WaitGroup счетчика
	wg.Done()
}

