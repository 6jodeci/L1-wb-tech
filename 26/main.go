// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
// Например: abcd — true || abCdefAaf — false || aabcd — false.
package main

import "fmt"

func UniqLetters(str string) (string, bool) {
	// этот будет перебирать все символы в строке до предпоследнего (включительно)
	// каждый из этих символов мы будем сравнивать с теми, которые идут после него
	for i := 0; i < len(str); i++ {
		// все элементы с индексом i будут сравниваться с элементами с индексом j
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return str, false
			}
		}
	}

	return str, true
}

func main() {
	fmt.Println(UniqLetters("abcd"))
	fmt.Println(UniqLetters("abCdefAaf"))
	fmt.Println(UniqLetters("aabcd"))
}
