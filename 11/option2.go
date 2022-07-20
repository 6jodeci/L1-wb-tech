// Реализовать пересечение двух неупорядоченных множеств.
// ВАРИАНТ 2
package main

import "fmt"

func main() {
	first := []string{"egor", "ivan", "dmitry", "sasha"} // первое множество
	second := []string{"egor", "sasha"}                  // второе множество

	fmt.Println("Пересечения: ", intersection(first, second))
	fmt.Println("Пересечения: ", intersection(second, first))
}

func intersection(first, second []string) []string {
	out := []string{}
	bucket := map[string]bool{}

	for _, i := range first {
		for _, j := range second {
			if i == j && !bucket[i] { // если значения массива 1 == значению маасива 2 и ещё не записано в баскет, записываем
				out = append(out, i) // аппендим в массив
				bucket[i] = true     // выходим из цикла если нашли все пересечения
			}
		}
	}

	return out
}
