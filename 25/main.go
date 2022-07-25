// Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"
)

func myOwnSleep(d int) {
	// After () используется для ожидания
	<-time.After(time.Second * time.Duration(d))
}

func main() {
	fmt.Println("Hello")
	myOwnSleep(3)
	fmt.Println("Sleep func!")
}
