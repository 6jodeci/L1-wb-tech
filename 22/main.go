// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// big.NewFloat для корректной работы с большими числами
	a := big.NewFloat(10000000)
	b := big.NewFloat(20000000)

	addition := new(big.Float).Add(a, b)
	multiplication := new(big.Float).Mul(a, b)

	subtraction := new(big.Float).Sub(a, b)
	division := new(big.Float).Quo(a, b)

	// Вывод результата
	// Глагол 1,f чтобы корректно отобразить результат
	fmt.Printf("Результат сложения: %1.f\n", addition)
	fmt.Printf("Результат умножения: %1.f\n", multiplication)

	fmt.Printf("Результат вычитаня: %1.f\n", subtraction)
	fmt.Println("Результат деления:", division)
}
