// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// структура Counter для хранения значения на которое увеличить счетчик
type Counter struct {
	value int64
}

func (count *Counter) inc(wg *sync.WaitGroup) {
	// уменьшаем счетчик wg на 1
	defer wg.Done()
	// атомарно добавляет единичку и возвращает значение
	atomic.AddInt64(&count.value, 1)
}

// вводим число и число на которое увиличить
func main() {
	wg := &sync.WaitGroup{}
	var value int
	fmt.Print("Введите число: ")
	fmt.Scan(&value)
	// увеличение счетчика (n раз)
	var n int64
	fmt.Print("На сколько увеличить: ")
	fmt.Scan(&n)
	val := Counter{n}
	for i := 0; i < value; i++ {
		// увеличиваем wg счетчик на 1
		wg.Add(1)
		// конкурентно запускаем функцию инкремента значения
		go val.inc(wg)
	}
	// ждем пока функция отработает
	wg.Wait()
	fmt.Println("Результат:", val)
}
