package concurrency

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type urlTracker struct {
	mu     sync.Mutex
	urlMap map[string]int
}

func newUrlTracker() urlTracker {
	return urlTracker{
		urlMap: make(map[string]int),
	}
}

func (t *urlTracker) isNewUrl(url string) bool {
	t.mu.Lock()
    defer t.mu.Unlock()
	newUrl := true
	if t.urlMap[url] > 0 {
		newUrl = false
	}
	t.urlMap[url]++
	return newUrl
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	tracker := newUrlTracker()
	c := make(chan string)

	var crawler func(url string, depth int, ret chan string)
	crawler = func(url string, depth int, ret chan string) {
		defer close(ret)

		// don't go deeper
		if depth <= 0 || !tracker.isNewUrl(url) {
			return
		}
		body, urls, err := fetcher.Fetch(url)

		// this level
		// error
		if err != nil {
			ret <- fmt.Sprint(err)
			return
		}
		// success
		ret <- fmt.Sprintf("found: %s %q", url, body)

		// downstream
		results := make([]chan string, len(urls))
		for i, u := range urls {
			results[i] = make(chan string)
            go crawler(u, depth-1, results[i])
		}
		// collect results from downstream
		for i := range results {
			for res := range results[i] {
				ret <- res
			}
		}
	}

	go crawler(url, depth, c)

	for entry := range c {
		fmt.Println(entry)
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func TestCralwer() {
	Crawl("https://golang.org/", 4, fetcher)
}
