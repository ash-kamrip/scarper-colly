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
		// magnet found
		metadata.Find("a").Each(func(i int, s *goquery.Selection) {
			if magnet, exists := s.Attr("href"); exists {
				if strings.Contains(magnet, "magnet") {
					fmt.Println(magnet)
				}
			}
		})
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
func movieFromString(data string, magnet string) Movie {
	elements := strings.Split(data, "\n")
	// Need to find a better way to do the parsing. This is fragile.
	return Movie{
		Name:     elements[1],
		Uploaded: elements[4],
		Magnet:   magnet,
		Size:     "formatSize1337x(elements[5])",
		Seeds:    elements[2],
		Uploader: elements[6],
	}
}

func instantTest(url string) ([]string, []Movie) {

	var links []string
	var movielist []Movie

	c := colly.NewCollector()

	// scrape the links from the whole page
	c.OnHTML("tr", func(h *colly.HTMLElement) {

		doc := h.DOM
		anchortags := doc.Find("a").Siblings()

		if attr, exists := anchortags.Attr("href"); exists {
			if strings.Contains(attr, "/torrent/") && !strings.HasPrefix(attr, "http://") {
				link := "https://1337x.to" + attr
				links = append(links, link)
			}
		}

		newmovie := movieFromString(doc.Text(), "magnet")
		movielist = append(movielist, newmovie)

	})
	// deep - 2nd
	c.OnHTML("div.col-9 div.box-info", func(h *colly.HTMLElement) {

		metadata := h.DOM

		movieName := metadata.Find("div.box-info-heading h1").Text()
		// fmt.Println(metadata.Find("div.box-info-heading h1").Text())
		// magnet found
		metadata.Find("a").Each(func(i int, s *goquery.Selection) {
			if magnet, exists := s.Attr("href"); exists {
				if strings.Contains(magnet, "magnet") {
					for _, movie := range movielist {
						fmt.Println(strings.TrimSpace(movie.Name))
						fmt.Println(strings.TrimSpace(movieName))
						if strings.TrimSpace(movie.Name) == strings.TrimSpace(movieName) {
							fmt.Println("strings matches now bitches")
							movie.Magnet = magnet
						} else {
							fmt.Println("ahhh fuck the name doesn't match")
						}

					}
				}
			}
		})
	})
	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
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
	return links, movielist[1:]

}
func main() {

	url := "https://1337x.to/search/spiderman/1/"
	links, movies := instantTest(url)

	fmt.Println("links to scrape")
	fmt.Println("---------------------------------------")
	for _, link := range links {
		fmt.Println(link)
	}
	fmt.Println("---------------------------------------")

	fmt.Println("movies : ")
	fmt.Println("---------------------------------------")
	for _, movie := range movies {
		fmt.Println("-------movie------")
		fmt.Println(movie.Name)
		fmt.Println(movie.Uploaded)
		fmt.Println(movie.Seeds)
		fmt.Println(movie.Uploader)
		fmt.Println(movie.Size)
		fmt.Println(movie.Magnet)
		fmt.Println("-------------")
	}
	fmt.Println("---------------------------------------")
	// url := "https://1337x.to/torrent/5137042/SpiderMan-No-Way-Home-2021-1080p-HD-TS-V3-Line-Audio-x264-AAC/"
	// scrapev2(url)
}
