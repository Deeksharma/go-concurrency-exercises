package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(message string, ch chan<- string) {
	// send message on ch1
	ch <- message
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1
	// send it on ch2

	ch2 <- <-ch1
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine genMsg and relayMsg
	go genMsg("hello", ch1)
	go relayMsg(ch1, ch2)
	// recv message on ch2
	msg := <-ch2
	fmt.Println(msg)
	close(ch1)
	close(ch2)
}
