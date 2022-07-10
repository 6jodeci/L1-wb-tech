// Реализовать все возможные способы остановки выполнения горутины. 
// ВАРИНАТ 2
package main

import "sync"

func main() {
	var wg sync.WaitGroup
	// увеличение счетчика на 1
	wg.Add(1)

	ch := make(chan int)
	go func() {
		for {
			foo, ok := <-ch
			if !ok {
				println("gorutine finished")
				// уменьшение счетчика на 1
				wg.Done()
				return
			}
			println(foo)
		}
	}()
	// запись 10 значений в канал
	for i := 0; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	// ждать завершений горутины
	wg.Wait()
}
