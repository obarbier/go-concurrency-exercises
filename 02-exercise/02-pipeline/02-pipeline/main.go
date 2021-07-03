// generator() -> square() -> print

package main

import (
	"fmt"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return merge(out)
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, c := range cs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for i := range c {
				out <- i
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := generator(2, 3, 7, 8, 9, 13, 221)

	// fan out square stage to run two instances.
	ch2a := square(in)
	ch2b := square(in)

	// TO fan in the results of square stages.
	for c := range merge(ch2a, ch2b) {
		fmt.Println(c)
	}

}
