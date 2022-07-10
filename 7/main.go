// Реализовать конкурентную запись данных в map.
package main

import "fmt"

type KeyValue struct {
	Key string
	Value string
}

// запись данных для Map в канал
func writeInMap (ch chan KeyValue) {
	ch <- KeyValue{"Petr", "Petrov"}
	ch <- KeyValue{"Egor", "Egorov"}
	ch <- KeyValue{"Ivan", "Ivanov"}
	ch <- KeyValue{"Dima", "Dimov"}
	ch <- KeyValue{"Sasha", "Sashev"}
	// закрытие канала
	close(ch)
}

func main() {
	// создание канала
	ch := make(chan KeyValue)
	// создание map
	mp := make(map[string]string)

	go writeInMap(ch)
	// перебор всех записаных в канал данных и запись в map
	for vl := range ch {
		mp[vl.Key] = vl.Value
	} 
	fmt.Println(mp)
}