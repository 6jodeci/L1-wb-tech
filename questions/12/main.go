package main

import "fmt"

func main() {
	n := 1
	if true {
		n := 1 // новая локальная переменная || if n = 1 || result = 2 (n++)
		n++
	}
	fmt.Println(n) // 1
}
