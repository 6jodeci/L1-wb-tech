// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "three"}

	writeMap := make(map[string]struct{})

	for _, uniq := range arr {
		// составной литерал 
		// https://go.dev/ref/spec#Composite_literals
		writeMap[uniq] = struct{}{}
	}
	fmt.Println(writeMap) // map[cat:{} dog:{} three:{}]
	// вытаскиваем все элементы из мапы
	for uniq := range writeMap {
		fmt.Println("Множество:", uniq)
	}
}