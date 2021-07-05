package main

import (
	"context"
	"fmt"
	"runtime"
)

func main() {

	// generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.
	generator := func(c context.Context) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			n := 0
			for {
				select {
				case out <- n:
					n++
				case <-c.Done():
					return

				}
			}
		}()
		return out
	}

	// Create a context that is cancellable.
	ctx, cancel := context.WithCancel(context.Background())

	count := 1
	for c := range generator(ctx) {
		fmt.Println(c)
		count++

		if count == 5 {
			cancel()
		}
	}

	fmt.Printf("number of running goroutine: %d", runtime.NumGoroutine())
}
