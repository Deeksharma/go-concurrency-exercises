package main

import "fmt"

func main() {
	//TODO: create channel owner goroutine which return channel and
	// writes data into channel and
	// closes the channel when done.
	owner := func() <-chan string { // channel owner function, responsible for initializing channels, writing in channel
		ch := make(chan string)
		go func() {
			defer close(ch)
			ch <- "yo yo 1"
			ch <- "yo yo 2"
			ch <- "yo yo 3"
			ch <- "yo yo 4"
			ch <- "yo yo 5"
		}()

		return ch
	}

	consumer := func(ch <-chan string) {
		// read values from channel
		for v := range ch {
			fmt.Printf("Received: %v\n", v)
		}
		fmt.Println("Done receiving!")
	}

	ch := owner()
	consumer(ch)
}
