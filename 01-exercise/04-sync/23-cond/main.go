package main

import (
	"fmt"
	"sync"
)

var sharedRsc sync.Map

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		val, _ := sharedRsc.Load("rsc1")
		fmt.Println(val)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		val, _ := sharedRsc.Load("rsc2")
		fmt.Println(val)
	}()
	// writes changes to sharedRsc
	sharedRsc.LoadOrStore("rsc1", "foo")
	sharedRsc.LoadOrStore("rsc2", "bar")
	wg.Wait()
}
