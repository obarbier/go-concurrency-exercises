package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

var fetched map[string]bool
var mapurl = make(map[string]bool)
var wg sync.WaitGroup
var mu sync.Mutex

func santitizeList(unsanitizedList []string) []string {
	var res []string
	checked := make(map[string]bool)
	for _, url := range unsanitizedList {
		checked[url] = true
	}

	for key := range checked {
		res = append(res, key)
	}

	return res

}

// Crawl uses findLinks to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int) {
	// Fetch URLs in parallel.
	if depth < 0 {
		return
	}
	mu.Lock()
	fetched[url] = true
	mu.Unlock()
	urls, err := findLinks(url)
	if err != nil {
		// fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", url)
	for _, u := range santitizeList(urls) {
		mu.Lock()
		if !fetched[u] {
			wg.Add(1)
			go func(u string, depth int) {
				defer wg.Done()
				Crawl(u, depth-1)
			}(u, depth)
		}
		mu.Unlock()
	}
	return
}

func main() {
	fetched = make(map[string]bool)
	now := time.Now()
	Crawl("http://andcloud.io", 2)
	wg.Wait()
	fmt.Println("time taken:", time.Since(now))
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// visit appends to links each link found in n, and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
