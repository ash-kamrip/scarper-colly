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

func scrapev2(url string) (mvlist []Movie) {

	c := colly.NewCollector()
	// selection criteria

	// scrape the links from the whole page
	c.OnHTML("a[href]", func(h *colly.HTMLElement) {

		doc := h.DOM
		if attr, exists := doc.Attr("href"); exists {
			if strings.Contains(attr, "/torrent/") && !strings.HasPrefix(attr, "http://") {
				link := "https://1337x.to" + attr
				c.Visit(link)
			}
		}

	})

	c.OnHTML("div.col-9 div.box-info", func(h *colly.HTMLElement) {

		metadata := h.DOM
		// fmt.Println(metadata.Find("div.box-info-heading h1").Text())

		// movieName
		metadata.Find("h1").Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Text())

		})
		// magnet found
		metadata.Find("a").Each(func(i int, s *goquery.Selection) {
			if magnet, exists := s.Attr("href"); exists {
				if strings.Contains(magnet, "magnet") {
					fmt.Println(magnet)
				}
			}
		})
		// seeds
		//

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

	return

}
func main() {

	url := "https://1337x.to/torrent/5137042/SpiderMan-No-Way-Home-2021-1080p-HD-TS-V3-Line-Audio-x264-AAC/"
	scrapev2(url)
}
