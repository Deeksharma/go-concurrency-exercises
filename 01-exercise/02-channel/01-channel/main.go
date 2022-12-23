package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 0)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)
	r := <-ch
	// TODO: get the value computed from goroutine
	fmt.Printf("computed value %v\n", r)
}
