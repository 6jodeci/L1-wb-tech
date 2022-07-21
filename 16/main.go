// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Создание рандомного слайса с n чисел (можно поменять количество чисел в слайсе)
	randomArray := createRadnomArray(10)
	// вывод слайса до сортировки
	fmt.Println("Unsorted slice:\n", randomArray)

	quicksort(randomArray)

	// вывод слайса после сортировки
	fmt.Println("\nSorted slice:\n", randomArray)
}

// createRadnomArray создает рандомный слайс с n чисел
func createRadnomArray(size int) []int {
	slice := make([]int, size, size)
	// позволяет сгенерировать рандомное значение
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		// генерирует рандомные числа для массива в пределах +-1000
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

// quicksort выполняет саму сортировку
func quicksort(s []int) []int {
	// если длина слайса меньше 2х - сразу выводим результат - сортировка не нужна
	if len(s) < 2 {
		return s
	}

	// этап разбиения
	left, right := 0, len(s)-1 // low(left) всегда = 0 || hight(right) всегда = длине массива -1
	center := rand.Int() % len(s) // опорный элемент
	s[center], s[right] = s[right], s[center]

	// по очереди выбираем элементы и сравниваем их с опорным (в левой части оказываются все элементы меньше опорного)
	// в правой части все элементы больше опорного
	for i := range s {
		if s[i] < s[right] {
			s[left], s[i] = s[i], s[left]
			left++
		}
	}
	s[left], s[right] = s[right], s[left]
	// рекурсивный вызов функции для левых и правых элементов
	quicksort((s[:left]))
	quicksort(s[left+1:])
	// возвращаем отсортированный слайс
	return s
}
