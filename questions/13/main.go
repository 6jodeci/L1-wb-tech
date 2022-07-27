package main

import "fmt"

func someAction(v []int8, b int8) { // int8
	v[0] = 100 // нулевой (первый) элемент == 100 || можно задать любой порядковый номер в слайсе
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5} // int8 // slice lenght == 5
	someAction(a, 6)              //
	fmt.Println(a)                // [100 2 3 4 5]
}
