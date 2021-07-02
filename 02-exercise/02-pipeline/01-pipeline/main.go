package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	ch1 := make(chan int)

	go func() {
		for _, num := range nums {
			ch1 <- num
		}
		close(ch1)
	}()
	return ch1
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	ch2 := make(chan int)

	go func() {
		for val := range in {
			ch2 <- val * val
		}
		close(ch2)
	}()

	return ch2
}

func main() {
	// set up the pipeline

	for i := range square(generator(1, 2, 3, 4, 5, 6, 7)) {
		fmt.Println(i)

	}
	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

}
