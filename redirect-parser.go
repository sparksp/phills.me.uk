package main

import (
	"fmt"
	"os"
	"path"
	"net/url"

	"golang.org/x/net/html"
)

type Redirect struct {
	from string
	to   string
}

func (r Redirect) String() string {
	return fmt.Sprintf("/%s  %s", r.from, r.to)
}

func stripDomain(href string) string {
	u, err := url.Parse(href)
	if err != nil {
		return ""
	}
	return u.Path
}

func getCanonical(t html.Token) (href string, ok bool) {
	for _, a := range t.Attr {
		switch a.Key {
		case "rel":
			if a.Val == "canonical" {
				ok = true
			} else {
				return "", false
			}
		case "href":
			href = stripDomain(a.Val)
		}
	}
	return
}

func stripRoot(filename string, root string) string {
	if len(filename) < len(root) {
		return filename
	}
	if filename[:len(root)] != root {
		return filename
	}
	return filename[len(root):]
}

func removePathIndex(filename string) string {
	if path.Base(filename) != "index.html" {
		return filename
	}
	return path.Dir(filename) + "/"
}

func crawl(filename string, chRedirect chan Redirect, chFinished chan bool) {
	defer func() {
		chFinished <- true
	}()
	var (
		from = stripRoot(removePathIndex(filename), "public/")
		to   string
	)
	defer func() {
		if from != "" && to != "" {
			chRedirect <- Redirect{from, to}
		}
	}()

	b, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + filename + "\"")
		return
	}
	z := html.NewTokenizer(b)

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return
		case html.SelfClosingTagToken:
			t := z.Token()

			switch t.Data {
			case "link":
				if url, ok := getCanonical(t); ok {
					to = url
				}
			}
		}
	}
}

func main() {
	chRedirect := make(chan Redirect)
	chFinished := make(chan bool)

	searchFiles := os.Args[1:]

	for _, filename := range searchFiles {
		go crawl(filename, chRedirect, chFinished)
	}

	for c := 0; c < len(searchFiles); {
		select {
		case redirect := <-chRedirect:
			fmt.Println(redirect)
		case <-chFinished:
			c++
		}
	}
}
