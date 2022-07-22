// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.
package main

import "fmt"

// ReverseStrint переворачивает строку
func ReverseString(s string) (result string) {
	for _, val := range s {
		// result = Текущий элемент + предыдущий result
		result = string(val) + result
	}
	return
}

func main() {
	var enter string
	fmt.Print("Enter string: ")
	fmt.Scan(&enter)
	fmt.Println("Reversed string:", ReverseString(enter))
}
