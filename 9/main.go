// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go in(array, ch1)
	go out(ch2, ch1)

	// вывод числе из канала 2 (val(ch1) * 2)
	for val := range ch2 {
		fmt.Println("val * 2: ", val)
	}
}

// запись чисел из массива в первый канал
func in(array []int, ch1 chan int) {
	for _, val := range array {
		ch1 <- val
	}
	close(ch1)
}

// проходим по всем числам из первого канала и умножаем на 2
func out(ch2 chan int, ch1 chan int) {
	for val := range ch1 {
		ch2 <- val * 2
	}
	close(ch2)
}
