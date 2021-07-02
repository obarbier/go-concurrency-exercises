package main

import "fmt"

func main() {
	c :=make(chan int)
	go func(a , b int , c chan int) {
		c <- a + b
	}(1, 2, c)
	fmt.Printf("computed value %v\n", <-c)
}
