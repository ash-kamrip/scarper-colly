package main

import (
	"fmt"
	"log"

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

	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

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
	movielist := getinfo(url)

	// print out the results
	for _, movie := range movielist {
		fmt.Println(movie)
	}
}
