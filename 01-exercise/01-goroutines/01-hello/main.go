package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("direct call version1")

	// goroutine with anonymous function
	go func() {
		fun("direct call version2")
	}()

	// goroutine with function value call
	fv := fun

	go fv("direct call version3")
	// wait for goroutines to end
	time.Sleep(100 * time.Millisecond)

	fmt.Println("done..")
}
