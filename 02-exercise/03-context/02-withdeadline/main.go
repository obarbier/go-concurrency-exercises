package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.
	deadline := time.Now().Add(10 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)

			deadline, ok := ctx.Deadline()
			if ok { // dealine is set
				if deadline.Sub(time.Now().Add(50*time.Millisecond)) < 0 { // if the deadline wiill be less than 50 milisecond
					fmt.Println("Not enough time to process this request")
					return // return is important since the select statement can either do ch or ctx.Done()
				}

			}
			// Simulate work.
			time.Sleep(50 * time.Millisecond)

			// Report result.	
			select {
			case ch <- data{"123"}:
			case <-ctx.Done():
				return
			}

		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	ch := compute(ctx)
	d, ok := <-ch
	if ok { // if d is comming send operation
		fmt.Printf("work complete: %s\n", d)
	}

}
