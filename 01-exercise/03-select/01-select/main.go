package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	for {
		select {
		case m, ok := <-ch1:
			if !ok {
				break
			}
			fmt.Println("ch1: " + m)
		case m, ok := <-ch2:
			if !ok {
				break
			}
			fmt.Println("ch2: " + m)
		}
	}

	// TODO: multiplex recv on channel - ch1, ch2

}
