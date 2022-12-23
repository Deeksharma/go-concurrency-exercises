package main

import (
	"fmt"
)

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, val := range nums {
			ch <- val
		}
		close(ch)
	}()
	return ch
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range in {
			out <- val * val
		}
		close(out)
	}()
	return out
}

func main() {
	// set up the pipeline
	for val := range square(square(generator(1, 2, 3, 3, 5, 6))) {
		fmt.Println(val)
	}
	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

}
