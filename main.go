package main

import (
	"fmt"
	"log"
	"strings"

	colly "github.com/gocolly/colly"
)

func getinfo(url string) []string {

	var links []string
	var magnet string

	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	// selection criteria
	c.OnHTML("a[href]", func(h *colly.HTMLElement) {

		doc := h.DOM

		if attr, exists := doc.Attr("href"); exists {
			if strings.Contains(attr, "/torrent/") && !strings.HasPrefix(attr, "http://itorrents.org") && !strings.HasPrefix(attr, "http://btcache.me") {
				link := "https://1337x.to" + attr
				c.Visit(link)
			}
			if strings.Contains(attr, "magnet:?") {
				magnet = attr
			}
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished")
	})

	// tells the collector to start scraping
	c.Visit(url)

	return links

}

func main() {

	url := "https://1337x.to/search/spiderman/1/"

	// getinfo(url)
	links := getinfo(url)

	// print out the results
	for _, link := range links {
		fmt.Println(link)
	}
}
