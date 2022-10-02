package main

import (
	"fmt"
	"log"
	"strings"

	colly "github.com/gocolly/colly"
)

func getinfo(url string) []string {

	var links []string
	var magnetlinks []string

	c := colly.NewCollector(
		colly.MaxDepth(2),
	)

	// selection criteria
	c.OnHTML("a[href]", func(h *colly.HTMLElement) {

		doc := h.DOM

		if attr, exists := doc.Attr("href"); exists {
			if strings.Contains(attr, "/torrent") {
				link := "https://1337x.to" + attr
				links = append(links, link)
			}
			if strings.Contains(attr, "magnet") {
				magnetlinks = append(magnetlinks, attr)
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

	for _, link := range links {
		c.Visit(link)
	}

	return magnetlinks[:1]

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
