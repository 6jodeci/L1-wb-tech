// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

// ReverseWords переворачивает слова в строке
func ReverseWords(s string) (result string) {
	// сначала преобразуем строку в массив строк, где каждая запись является словом
	words := strings.Fields(s)
	// обратный цикл
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		// меняем слова местами
		words[i], words[j] = words[j], words[i]
	}
	// соединяем все строки
	return strings.Join(words, " ")
}

func main() {
	// можно поменять строку
	enter := "snow dog sun"
	fmt.Println("Words before reverse: ", enter)
	// результат
	fmt.Println("Words after reverse:", ReverseWords(enter))
}
