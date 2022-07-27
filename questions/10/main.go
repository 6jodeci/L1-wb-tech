package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a // p ссылается на a == 1
	)
	fmt.Println(*p) // p == 1
	update(p)
	fmt.Println(*p) // p также = 1
}
