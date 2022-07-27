package main

import "fmt"

func main() {
	slice := []string{"a", "a"} // слайс main функции || result = [a a] || новая локальная переменная

	func(slice []string) { // новая локальная переменная
		slice = append(slice, "a") // append "a" || result = [a]
		slice[0] = "b"             // append "b" в нулевой индекс || result = [b a]
		slice[1] = "b"             // append "b" в первый индекс || result = [b b a]
		fmt.Print(slice)           // локальный слайс в созданой функции result [b b a]
	}(slice)
	fmt.Print(slice) // result [a a]
	// из-за использования print (перенос строки не обозначен \n reslut = [b b a] [a a])
}
