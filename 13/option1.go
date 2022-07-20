// Поменять местами два числа без создания временной переменной.
// ВАРИАНТ 1
package main

import "fmt"

func main() {
	var(
		num1 = 10
		num2 = 20
	)
	fmt.Println("Числа до:", num1, "  ", num2)
	num1 = num1 + num2 // num1 (10 + 20) == 30
	num2 = num1 - num2 // num2 (30 - 20) == 10
	num1 = num1 - num2 // num1 (30 - 10) == 20
	fmt.Println("Числа после", num1,"  ", num2)
}
