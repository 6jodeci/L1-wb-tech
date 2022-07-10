// Реализовать все возможные способы остановки выполнения горутины. 
// ВАРИАНТ 1
package main

import "fmt"

func main() {
	// передача пустой структуры (0)
	quit := make(chan struct{})
    // или передаем false (0)
    // quit := make(chan bool)
go func() {
    for {
        select {
        case <-quit:
            return
        default:
            
        }
    }
}()
	fmt.Println("gorutine finished")
	// закрыте канала
	close(quit)
}