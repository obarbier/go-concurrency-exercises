package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex
	var murw sync.RWMutex
	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	//TODO: implement concurrent read.
	// allow multiple reads, writes holds the lock exclusively.
	for i := 0; i<6; i++ {
		go func() {
			murw.RLock()
			defer murw.RUnlock()
			fmt.Println(balance)
		}()
	}

	wg.Wait()
	fmt.Println(balance)
}
