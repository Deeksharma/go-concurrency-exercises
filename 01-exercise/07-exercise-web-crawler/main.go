package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

var fetched map[string]bool

// Crawl uses findLinks to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int) {
	// TODO: Fetch URLs in parallel.

	if depth < 0 {
		return
	}
	urls, err := findLinks(url)
	if err != nil {
		// fmt.Println(err)
		return
	}
	fmt.Printf("found: %s\n", url)
	fetched[url] = true
	for _, u := range urls {
		if !fetched[u] {
			Crawl(u, depth-1)
		}
	}
	return
}

func main() {
	fetched = make(map[string]bool)
	now := time.Now()
	Crawl("http://andcloud.io", 2)
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

/*
In the crawl function, it is calling function find links and passing it an url.

And the find links function, we do a http get on the url, we check if the status is okay, if not,

we return an error status.

We do a http parse on the response body.

And we pass the html document that we get to the visit function.

in the visit function, we find the links in the document.

And we create a slice of the links.

And here we are returning the slice of the urls.

We come back to the crawl function.

We receive the slice of the urls.

We print the url, which we currently crawled.

And we are maintaining a map of the urls, so that we don't crawl the same url again.

We range over the urls and we recursively call the crawl function on each url till the specified depth.
*/
