package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print
var ch1, ch2 chan int

// generator - convertes a list of integers to a channel
func generator(nums ...int) {
	// ch1 = make(chan int)
	for num := range nums {
		ch1 <- num
	}
	close(ch1)

}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square() {
	// ch2 = make(chan int)
	for val := range ch1 {
		ch2 <- val * val
	}
	close(ch2)
}

func main() {
	ch1, ch2 = make(chan int), make(chan int)
	// set up the pipeline
	go generator(1, 2, 3, 4, 5, 6, 7)
	go square()
	for i := range ch2 {
		fmt.Println(i)

	}
	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

}
