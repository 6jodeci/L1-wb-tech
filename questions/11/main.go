package main

import (
	"fmt"
	"sync"
)

// default - начнет перебор значений в цикле и их вывод после чего улетит в deadlock не выйдя из цикла
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

// func main() {
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 5; i++ {
// 			fmt.Println(i)
// 		}
// 	}()
// 	wg.Wait()
// 	fmt.Println("exit")
// }
