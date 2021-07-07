package main

import "fmt"

// ByteCounter type
type ByteCounter int

// TODO: Implement Write method for ByteCounter
// to count the number of bytes written.
func (b *ByteCounter) Write(p []byte) (n int, err error) {
	n = len(p)
	*b = ByteCounter(n)
	return n, nil
}
func main() {
	var b ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
