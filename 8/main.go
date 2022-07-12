// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import "fmt"

func main() {
	// Вводим с клавиатуры число для которого хотим изменить биты
	var number int64 
	fmt.Print("Введите число у которого будем менять i-ый бит: ")
	fmt.Scan(&number)
	fmt.Printf("Введенное число: %d в двоичной системе: %b\n", number, number)
	// if true i = 1 || if false i = 0
	valBool := false
	// Выбор бита который будем менять
	valByte := 1
	if valBool {
		number |= 1 << valByte // если 1, установка с помощью логического ИЛИ
	} else {
		number &^= 1 << valByte // если 0, сброс с помощью логического ИЛИ НЕ
	}

	fmt.Printf("Введенное число: %d в двоичной системе после изменения: %b\n", number, number)
}
