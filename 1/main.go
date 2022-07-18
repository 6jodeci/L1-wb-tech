// Дана структура Human (с произвольным набором полей и методов). 
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)

package main 

import "fmt"

// Родительская
type Human struct {

}

// Наследование от родительской
type Action struct {
	Human
}

// Реализация метода для родительской 
func (h *Human) Code() {
	fmt.Println("Coding..")
}

// Реализация метода для родительской 
func (h *Human) Sleep() {
	fmt.Println("Sleeping..")
}

// Наследование метода
func main () {
	action := Action{}
	action.Code()
	action.Sleep() 
}
