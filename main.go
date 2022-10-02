package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	colly "github.com/gocolly/colly"
)

type Movie struct {
	Name     string
	Uploaded string
	Magnet   string //link webtorrent
	Size     string
	Seeds    string //as of new seeds are already sorted in descending order in 1337x.to
	Uploader string
}

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

			fmt.Println("movie name and magnet found")
			fmt.Println()
		}
	})

	c.OnHTML("div.col-9 div.box-info", func(h *colly.HTMLElement) {

		metadata := h.DOM
		// fmt.Println(metadata.Find("div.box-info-heading h1").Text())
		// magnet found
		metadata.Find("a").Each(func(i int, s *goquery.Selection) {
			if magnet, exists := s.Attr("href"); exists {
				if strings.Contains(magnet, "magnet") {
					fmt.Println(magnet)
				}
			}
		})
		// seeds
		// Uploaded
		// Size
		// Seeds
		// Uploader
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
