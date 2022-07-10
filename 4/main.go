// Реализовать постоянную запись данных в канал (главный поток). 
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
package main

import (
	"fmt"
	"time"
)

func main() {
	workerStart()
}

func workerStart() {
	var n int
	fmt.Println("введите количество горутин:")
	fmt.Scan(&n)

	// использование не буферизированного канала
	workerInput := make(chan int)

	// создаем горутины
	for i := 0; i < n; i++ {
		go worker(i, workerInput)
	}

	// запись в канал из потока
	for {
		workerInput <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

// воркеры
func worker(worker int, in <-chan int) {
	for {
		number := <-in
		fmt.Printf("goroutine #%d: value: %d\n", worker, number)
	}
}

