// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.
// ВАРИАНТ 2
package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := make(chan(int))
	i := reflect.ValueOf(t)
	switch i.Kind() {
	case reflect.String:
		fmt.Printf("type=sting || value=%s\n", i)
	case reflect.Int:
		fmt.Printf("type=integer || value=%v\n", i)
	case reflect.Bool:
		fmt.Printf("type=boolean || value=%v\n", i.Bool())
	case reflect.Float64:
		fmt.Printf("type=float64 || value=%v\n", i)
	case reflect.Chan:
		fmt.Printf("type=chan || value=%d\n", i.Interface())
	}
}
