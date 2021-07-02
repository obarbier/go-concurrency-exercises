package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 6)

	go func() {
		defer close(ch)

		// TODO: send all iterator values on channel without blocking
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			time.Sleep(3 * time.Millisecond)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
		time.Sleep(1 * time.Millisecond)
	}
}
