// Поменять местами два числа без создания временной переменной.
// ВАРИАНТ 2
package main

import "fmt"

func main() {
	var (
		num1 = 10
		num2 = 20
	)
	fmt.Println("Числа до:", num1, "  ", num2)
	num1 = num1 ^ num2
	num2 = num1 ^ num2
	num1 = num1 ^ num2
	fmt.Println("Числа после", num1, "  ", num2)
}
