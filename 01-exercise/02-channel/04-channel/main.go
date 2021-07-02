package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(out chan<- string, message string) {
	out <- message
}

func relayMsg(c2 chan<- string, c1 <-chan string) {
	c2 <- <-c1
}

func main() {
	// create ch1 and ch2
	var c1, c2 chan string
	c1 = make(chan string)
	c2 = make(chan string)
	// spine goroutine genMsg and relayMsg
	go genMsg(c1, "my message")

	// recv message on ch2
	go relayMsg(c2, c1)

	fmt.Println(<-c2)
}
