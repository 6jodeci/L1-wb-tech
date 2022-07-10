package main

import (
	"fmt"
	"time"
)

// время через которое завершится программа
const SECONDS_TO_STOP = 5

func main() {
	ch := make(chan int)
	go printRead(ch)
	go send(ch)
	// таймер на 5 секунд
	time.Sleep(time.Duration(SECONDS_TO_STOP *time.Second))
	// закрытие канала
	close(ch)
}

// вывод записанных значений в канал
func printRead(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

// запись в канал
func send(ch chan<- int) {
	i:= 0
	for {
		ch <- i
		i++
	}
}