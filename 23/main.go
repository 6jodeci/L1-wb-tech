// Удалить i-ый элемент из слайса.
package main

import "fmt"

func main() {
	slice := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Println("Исходный слайс: ", slice)

	var i int // i-ый элемент
	fmt.Print("Какой элемент удалить: ")
	fmt.Scan(&i)

	// сдвинуть a[i+1:] влево на один индекс.
	copy(slice[i:], slice[i+1:])

	// удалить последний элемент (записать нулевое значение)
	slice[len(slice)-1] = ""

	// уменьшить слайс
	slice = slice[:len(slice)-1]

	fmt.Println("Слайс после удаления i-того элемента: ", slice)
}
