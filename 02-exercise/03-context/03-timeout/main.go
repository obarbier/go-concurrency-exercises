package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// set a http client timeout

	duration := 10 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	process := func(ctx context.Context) <-chan *http.Response {
		out := make(chan *http.Response)

		go func() {
			defer close(out)
			req, err := http.NewRequest("GET", "https://andcloud.io", nil)
			if err != nil {
				log.Fatal(err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println("ERROR:", err)
				return
			}

			// Write the response to stdout.
			select {
			case <-ctx.Done():
				fmt.Println("Timeout")
				resp.Body.Close()
				return
			case out <- resp:
			}

		}()

		return out

	}

	c, ok := <-process(ctx)
	if ok {
		io.Copy(os.Stdout, c.Body)
		c.Body.Close()
	}
}
