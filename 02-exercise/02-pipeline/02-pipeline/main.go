// generator() -> square() -> print

package main

import "fmt"

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
	out1 := make(chan int)
	out2 := make(chan int)
	go func() {
		for n := range in {
			out1 <- n * n
		}
		close(out1)
	}()
	go func() {
		for n := range in {
			out2 <- n * n
		}
		close(out2)
	}()
	return merge(out1, out2)
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for _, c := range cs {
			for i := range c {
				out <- i
			}

		}
		close(out)
	}()
	return out
}

func main() {
	in := generator(2, 3, 7, 8, 9, 13, 221)

	// fan out square stage to run two instances.
	in2 := square(in)

	// TO fan in the results of square stages.
	for c := range in2 {
		fmt.Println(c)
	}

}
