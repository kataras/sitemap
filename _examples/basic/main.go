package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kataras/sitemap"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path))
	})

	sitemaps := sitemap.New("http://localhost:8080").
		URL(sitemap.URL{Loc: "/home", Links: []sitemap.Link{{Hreflang: "el", Href: "/el/home"}}}).
		URL(sitemap.URL{Loc: "/articles", LastMod: time.Now(), ChangeFreq: sitemap.Daily, Priority: 1}).
		URL(sitemap.URL{Loc: "/about"}).
		URL(sitemap.URL{Loc: "/rss", ChangeFreq: sitemap.Always}).
		Build()
		//
	for _, s := range sitemaps {
		fmt.Printf("Pos: %d\nPath: %s\nIsSitemapIndex: %v\n",
			s.Pos, s.Path, s.IsSitemapIndex)

		// http://localhost:8080/sitemap.xml
		// When more than 50000 URLs are registered then
		// we have more than one handler to register here
		// with paths like: http://localhost:8080/sitemap1.xml ...sitemap%d.xml.
		// Read more at: https://www.sitemaps.org/protocol.html
		mux.Handle(s.Path, s)
	}

	// http://localhost:8080/sitemap.xml
	http.ListenAndServe(":8080", mux)
}
