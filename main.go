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
func movieFromString(document *goquery.Selection) Movie {

	// elements := strings.Split(data, "\n")
	// // Need to find a better way to do the parsing. This is fragile.
	// return Movie{
	// 	Name:     elements[1],
	// 	Uploaded: elements[4],
	// 	Magnet:   magnet,
	// 	Size:     "formatSize1337x(elements[5])",
	// 	Seeds:    elements[2],
	// 	Uploader: elements[6],
	// }

	return Movie{
		Name:     document.Find("td.name > a:nth-child(2)").Text(),
		Uploaded: document.Find("td.coll-date").Text(),
		Size:     strings.Replace(document.Find("td.size").Text(), document.Find("td.seeds").Text(), "", -1),
		Seeds:    document.Find("td.seeds").Text(),
		Uploader: document.Find("td.coll-5").Text(),
	}
}

var Movies []Movie

func instantTest(url string) {

	var links []string
	movielist := make([]Movie, 0)

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

		newmovie := movieFromString(doc)
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
						// fmt.Println(strings.TrimSpace(movie.Name))
						// fmt.Println(strings.TrimSpace(movieName))
						if strings.TrimSpace(movie.Name) == strings.TrimSpace(movieName) {
							fmt.Println("the name matches")
							movie.Magnet = magnet
							Movies = append(Movies, movie)
						}
					}
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

	for _, link := range links {
		c.Visit(link)
	}

}
func main() {

	url := "https://1337x.to/search/spiderman/1/"
	instantTest(url)

	// fmt.Println("links to scrape")
	// fmt.Println("---------------------------------------")
	// for _, link := range links {
	// 	fmt.Println(link)
	// }
	// fmt.Println("---------------------------------------")

	fmt.Println("movies : ")
	fmt.Println("---------------------------------------")
	for _, movie := range Movies {
		fmt.Println("-------movie------")
		fmt.Println("Name of the movie : ", movie.Name)
		fmt.Println("uploade date : ", movie.Uploaded)
		fmt.Println("magnetlink : ", movie.Magnet)
		fmt.Println("file size: :", movie.Size)
		fmt.Println("seeds: ", movie.Seeds)
		fmt.Println("uploader : ", movie.Uploader)
		fmt.Println("-------------")
	}
	fmt.Println("---------------------------------------")
	// url := "https://1337x.to/torrent/5137042/SpiderMan-No-Way-Home-2021-1080p-HD-TS-V3-Line-Audio-x264-AAC/"
	// scrapev2(url)
}
