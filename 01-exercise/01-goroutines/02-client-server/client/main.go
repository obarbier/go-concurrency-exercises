package main

import (
	"io"
	"log"
	"net"
	"os"
	"fmt"
)

func main() {
	// Example from https://golang.org/pkg/net/
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
		fmt.Println("connection fail")
	}

	mustCopy( os.Stdout ,conn)

}

// mustCopy - utility function
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
