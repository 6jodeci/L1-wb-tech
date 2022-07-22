// Реализовать бинарный поиск встроенными методами языка.
package main

import (
	"fmt"
)

func main() {
	slice := []int{2, 0, 1, 5, 53, 34, 67, 76, 22, 11, 96, 54, 26}

	// число которое будем искать в слайсе
	n := 0
	// выозов функции бинарного поиска
	result, count := binarySearch(slice, n)
	if result < 0 {
		fmt.Printf("\nЧисло %v не найдено || Введите другое число!\n\n", n)
	} else {
		fmt.Printf("Число %v найдено || Номер в слайсе: %v || Выполненно: %v итерации \n\n", n, result, count)
	}
}

func binarySearch(slice []int, n int) (res int, count int) {
	// ищем центральный элемент слайса
	mid := len(slice) / 2
	switch {
	// если длина слайса = 0, значит элемента нет в слайсе, в
	case len(slice) == 0:
		res = -1
	// сравниваем центральный элемент с искомым, если меньше, ищем в левой части слайсла
	case slice[mid] > n:
		// рекурсивный вызов функции для левой части слайса
		res, count = binarySearch(slice[:mid], n)
		// если искомый элемент больше, ищем в правой части слайса
	case slice[mid] < n:
		// рекурсивный вызов функции для правой части слайса
		res, count = binarySearch(slice[mid+1:], n)
		// если элмент найден прибавляем единичку
		if res >= 0 {
			res += mid + 1
		}
	default:
		// если искомый элемент - это средний элемент
		res = mid
	}
	count++
	// выводим само число, и номер в слайсе
	return res, count
}
