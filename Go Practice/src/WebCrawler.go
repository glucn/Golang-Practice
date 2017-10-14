package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Crawler struct {
	visited map[string]bool
	mux     sync.Mutex
}

func (c *Crawler) visit(s string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.visited == nil {
		c.visited = make(map[string]bool)
	}

	if _, ok := c.visited[s]; ok {
		return
	} else {
		c.visited[s] = true
		return
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Crawler) Crawl(url string, depth int, fetcher Fetcher) {
	var wg sync.WaitGroup

	var crawler func(string, int)

	crawler = func(url string, depth int) {
		defer wg.Done()

		if depth <= 0 {
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			c.visit(fmt.Sprintf("%s\n", err))
			return
		}
		c.visit(fmt.Sprintf("found: %s %q\n", url, body))
		for _, u := range urls {
			wg.Add(1)
			go crawler(u, depth-1)
		}
	}

	wg.Add(1)
	crawler(url, depth)
	wg.Wait()
}

func main() {
	var c Crawler
	c.Crawl("http://golang.org/", 4, fetcher)
	for s := range c.visited {
		fmt.Printf(s)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}