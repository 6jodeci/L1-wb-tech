// Реализовать пересечение двух неупорядоченных множеств.
// ВАРИАНТ 1 
package main

import "fmt"

func intersection(one, two []int) []int {
	var m = make(map[int]uint) // создаем мапу для хранения чисел

	for i := range one { // пробегаясь по первому слайсу инкрементируем мапу
		m[one[i]]++ // увеличиваем мапу, т.к defaul uint value == 0
	}

	var result = make([]int, 0) // создаем массив с результатами

	for i := range two { // пробегаясь по второму - ищем пересечения
		if value, ok := m[two[i]]; ok {
			if value > 0 {
				m[two[i]]--
				result = append(result, two[i])
			} else {
				delete(m, two[i]) // прибираемся, так как ключ уже не нужен (== 0)
			}
		}
	}

	return result
}

func main() {
	fmt.Printf("Пересечения: %v\n", intersection([]int{23, 3, 1, 2}, []int{6, 2, 4, 23}))
}
