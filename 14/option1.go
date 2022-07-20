// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.
// ВАРИАНТ 1
package main

import "fmt"

func main() {
	var t interface{}
	// enter your type of variable
	t = 1.1

	switch t := t.(type) {

	default:
		fmt.Printf("type=string || value=%s\n", t) 
	case bool:
		fmt.Printf("type=boolean || value=%t\n", t) 
	case int:
		fmt.Printf("type=integer || value=%v\n", t) 
	case float64:
		fmt.Printf("type=float64 || value=%v\n", t) 
	case any:
		fmt.Printf("type=chan || value=%d\n", t)
	}

}
