// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
// ВАРИАНТ 2
package main

import (
	"fmt"
	"math"
	"sync"
)

var wg sync.WaitGroup

func main() {
	
	numbers := []float64{2, 4, 6, 8, 10}
	var sum float64

	for _, i := range numbers {
		// Увеличение WaitGroup на 1
		wg.Add(1)
		
		go func(number float64) {
			// Обнуление счетчика
			defer wg.Done()
			sum += math.Pow(number, 2)
		}(i)
		
	}
	// Ждем пока счетчик станет нулем
	wg.Wait()
	// Вывод суммы квадратов
	fmt.Println(sum)
}

