package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	cloudflarebp "github.com/DaRealFreak/cloudflare-bp-go"
	"github.com/PuerkitoBio/goquery"
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

func random(url string) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		// handle err
		log.Fatal(err)
	}

	client := &http.Client{}
	client.Transport = cloudflarebp.AddCloudFlareByPass(client.Transport)

	// specifying timeout settings
	ctx, cancel := context.WithTimeout(req.Context(), 10*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	res, err := client.Do(req)

	if err != nil {
		// handle err
		log.Fatal(err)
	}

	defer res.Body.Close()

	slurp, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%v\n%v\n%s\n", res, res.Request, slurp)

}

func learnJquery() {
	// var headings []string
	// var row []string
	// var rows [][]string
	data := `
<html>
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<title>Download spiderman Torrents | 1337x</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="/css/jquery-ui.css">
<link rel="stylesheet" href="/css/icons.css">
<link rel="stylesheet" href="/css/fancySelect.css">
<link rel="stylesheet" href="/css/style10.2022-10-01-09.css">
<link rel="shortcut icon" href="/favicon.ico">
<!--[if lt IE 9]><script src = "/js/html5shiv.js"></script><![endif]-->
<style id="antiClickjack">body{display:none !important;}</style>
<script type="text/javascript" id="antiClickjackJS">
if (self === top) {
var antiClickjack = document.getElementById("antiClickjack");
antiClickjack.parentNode.removeChild(antiClickjack);
} else {
top.location = self.location;
}
</script>
<script data-cfasync="false" src="//d2ers4gi7coxau.cloudfront.net/?gsred=949060"></script>
<script data-cfasync="false" src="/sw.js"></script>
</head>
<body>
<div class="mobile-menu"></div>
<div class="top-bar">
<div class="container">
<ul class="top-bar-nav">
<li><a href="/register">Register</a></li>
<li class="active"><a href="/login">Login</a></li>
</ul>
</div>
</div>
<header>
<div class="container">
<div class="clearfix">
<div class="logo"><a href="/home/"><img alt="logo" src="/images/logo.svg"></a></div>
<a href="#" class="navbar-menu"><span></span><span></span><span></span></a>
<div class="search-box">
<form id="search-form" method="get" action="/srch">
<input type="search" placeholder="Search for torrents.." value="spiderman" id="autocomplete" name="search" class="ui-autocomplete-input form-control" autocomplete="off">
<button type="submit" class="btn btn-search"><i class="flaticon-search"></i><span>Search</span></button>
</form>
</div>
</div>
<nav>
<ul class="main-navigation">
<li class="active"><a href="/home/" title="Go to Home">Home</a></li>
<li><a href="/upload" title="Upload Torrent File">Upload</a></li>
<li><a href="/rules" title="Rules">Rules</a></li>
<li><a href="/contact" title="Contact 1337x">Contact</a></li>
<li><a href="/about" title="About us">About us</a></li>
</ul>
</nav>
</div>
</header>
<main class="container">
<div class="row">
<aside class="col-3 pull-right">
<div class="list-box">
<h2>Browse torrents</h2>
<ul>
<li><a href="/trending" title="Trending Torrents"><i class="flaticon-trending"></i> Trending Torrents</a></li>
<li><a href="/movie-library/1/" title="Movie Library"><i class="flaticon-movie-library"></i> Movie library</a></li>
<li><a href="/series-library/a/1/" title="TV Library"><i class="flaticon-tv-library"></i> TV library</a></li>
<li><a href="/new-episodes/day/1/" title="New TV Episodes"><i class="flaticon-tv"></i> New TV Episodes</a></li>
<li><a href="/top-100" title="Top 100 Torrents"><i class="flaticon-top"></i> Top 100 Torrents</a></li>
<li><a href="/cat/Anime/1/" title="Anime"><i class="flaticon-ninja-portrait"></i> Anime</a></li>
<li><a href="/cat/Apps/1/" title="Applications"><i class="flaticon-apps"></i> Applications</a></li>
<li><a href="/cat/Documentaries/1/" title="Documentaries"><i class="flaticon-documentary"></i> Documentaries</a></li>
<li><a href="/cat/Games/1/" title="Games"><i class="flaticon-games"></i> Games</a></li>
<li><a href="/cat/Movies/1/" title="Movies"><i class="flaticon-movies"></i> Movies</a></li>
<li><a href="/cat/Music/1/" title="Movies"><i class="flaticon-music"></i> Music</a></li>
<li><a href="/cat/Other/1/" title="Others"><i class="flaticon-other"></i> Other</a></li>
<li><a href="/cat/TV/1/" title="Television"><i class="flaticon-tv"></i> Television</a></li>
<li><a href="/cat/XXX/1/" title="XXX"><i class="flaticon-xxx"></i> XXX</a></li>
</ul>
</div>
<div class="list-box hidden-sm">
<h2>1337x Links</h2>
<ul class="list">
<li><a target="_blank" href="https://www.1337x-forum.eu/"> 1337x Forum</a></li>
<li><a target="_blank" href="https://chat.1337x.to"> 1337x Chat</a></li>
<li><a target="_blank" href="https://novastream.to">NovaStream</a></li>
<li><a target="_blank" href="https://njal.la"> Njalla</a></li>
<li><a target="_blank" href="https://www.limetorrents.cc"> Limetorrents</a></li>
<li><a target="_blank" href="https://www.torrentfunk.com"> TorrentFunk</a></li>
<li><a target="_blank" href="https://www.torlock.com"> Torlock</a></li>
</ul>
</div>
</aside>
<div class="col-9 page-content search-page">
<div class="box-info">
<center>
<div><script data-cfasync="false" async type="text/javascript" src="//untrendenam.com/t1qHqmYLNrj/55830"></script></div>
</center>
<br> <div class="box-info-heading clearfix">
<h1> Searching for: <span>spiderman</span> in:</h1>
<div class="box-info-left sort-by-box">
<form action="#">
<select class="select" onChange="top.location.href=this.options[this.selectedIndex].value;">
<option value="#" selected="selected">All Categories</option>
<option value="/category-search/spiderman/Movies/1/">Movies Only</option>
<option value="/category-search/spiderman/TV/1/">TV Only</option>
<option value="/category-search/spiderman/Games/1/">Games Only</option>
<option value="/category-search/spiderman/Music/1/">Music Only</option>
<option value="/category-search/spiderman/Apps/1/">Applications Only</option>
<option value="/category-search/spiderman/Documentaries/1/">Documentaries Only</option>
<option value="/category-search/spiderman/Anime/1/">Anime Only</option>
<option value="/category-search/spiderman/Other/1/">Other Only</option>
<option value="/category-search/spiderman/XXX/1/">XXX Only</option>
</select>
</form>
</div>
<div class="box-info-right sort-by-box">
<form action="#">
<select class="select" onChange="top.location.href=this.options[this.selectedIndex].value;">
<option value="#" selected="selected">Sort by...</option>
<option value="/sort-search/spiderman/time/desc/1/">Sort by Time</option>
<option value="/sort-search/spiderman/size/desc/1/">Sort by Size</option>
<option value="/sort-search/spiderman/seeders/desc/1/">Sort by Seeders</option>
<option value="/sort-search/spiderman/leechers/desc/1/">Sort by Leechers</option>
</select>
</form>
</div>
</div>
<div class="box-info-detail inner-table">
<div class="table-list-wrap">
<table class="table-list table table-responsive table-striped">
<thead>
<tr>
<th class="coll-1 name">name</th>
<th class="coll-2">se</th>
<th class="coll-3">le</th>
<th class="coll-date">time</th>
<th class="coll-4"><span class="size">size</span> <span class="info">info</span></th>
<th class="coll-5">uploader</th>
</tr>
</thead>
<tbody>
<tr>
<td class="coll-1 name"><a href="/sub/54/0/" class="icon"><i class="flaticon-h264"></i></a><a href="/torrent/5137042/SpiderMan-No-Way-Home-2021-1080p-HD-TS-V3-Line-Audio-x264-AAC/">SpiderMan No Way Home 2021 1080p HD-TS V3 Line Audio x264 AAC</a><span class="comments"><i class="flaticon-message"></i>12</span></td>
<td class="coll-2 seeds">2250</td>
<td class="coll-3 leeches">256</td>
<td class="coll-date">Feb. 2nd '22</td>
<td class="coll-4 size mob-vip">2.5 GB<span class="seeds">2250</span></td>
<td class="coll-5 vip"><a href="/user/mazemaze16/">mazemaze16</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/54/0/" class="icon"><i class="flaticon-h264"></i></a><a href="/torrent/3860369/Spider-Man-Far-From-Home-2019-English-720p-CAMRip-NO-WATERMARKS-MP3-x264-HC-HD-Web-Movies/">Spider-Man Far From Home (2019) English 720p CAMRip [NO WATERMARKS] MP3 x264 HC [HD Web Movies]</a><span class="comments"><i class="flaticon-message"></i>42</span></td>
<td class="coll-2 seeds">1176</td>
<td class="coll-3 leeches">246</td>
<td class="coll-date">Jul. 2nd '19</td>
<td class="coll-4 size mob-uploader">1.1 GB<span class="seeds">1176</span></td>
<td class="coll-5 uploader"><a href="/user/Suryadipta/">Suryadipta</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/42/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/5178889/Spider-Man-No-Way-Home-2021-1080p-BRRip-x264-ProLover/">Spider-Man : No Way Home (2021) 1080p BRRip x264 - ProLover</a><span class="comments"><i class="flaticon-message"></i>6</span></td>
<td class="coll-2 seeds">624</td>
<td class="coll-3 leeches">324</td>
<td class="coll-date">Mar. 11th '22</td>
<td class="coll-4 size mob-uploader">1.4 GB<span class="seeds">624</span></td>
<td class="coll-5 uploader"><a href="/user/prolover/">prolover</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/54/0/" class="icon"><i class="flaticon-h264"></i></a><a href="/torrent/5135809/SpiderMan-No-Way-Home-2021-720p-HD-TS-V3-Line-Audio-x264-AAC/">SpiderMan No Way Home 2021 720p HD-TS V3 Line Audio x264 AAC</a><span class="comments"><i class="flaticon-message"></i>1</span></td>
<td class="coll-2 seeds">442</td>
<td class="coll-3 leeches">38</td>
<td class="coll-date">Feb. 1st '22</td>
<td class="coll-4 size mob-vip">1.2 GB<span class="seeds">442</span></td>
<td class="coll-5 vip"><a href="/user/mazemaze16/">mazemaze16</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/70/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/5179575/Spider-Man-No-Way-Home-2022-1080p-Bluray-10bit-DTS-HD-MA-5-1-x265-HashMiner/">Spider-Man.No.Way.Home.2022.1080p.Bluray.10bit.DTS-HD.MA.5.1.x265.[HashMiner]</a><span class="comments"><i class="flaticon-message"></i>12</span></td>
<td class="coll-2 seeds">422</td>
<td class="coll-3 leeches">116</td>
<td class="coll-date">Mar. 12th '22</td>
<td class="coll-4 size mob-uploader">8.2 GB<span class="seeds">422</span></td>
<td class="coll-5 uploader"><a href="/user/HashMiner/">HashMiner</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/42/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/4006899/Spider-Man-Far-from-Home-2019-720p-BluRay-x264-NeZu/">Spider-Man.Far.from.Home.2019.720p.BluRay.x264-NeZu</a><span class="comments"><i class="flaticon-message"></i>5</span></td>
<td class="coll-2 seeds">283</td>
<td class="coll-3 leeches">136</td>
<td class="coll-date">Sep. 19th '19</td>
<td class="coll-4 size mob-vip">1.4 GB<span class="seeds">283</span></td>
<td class="coll-5 vip"><a href="/user/NeZu/">NeZu</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/54/0/" class="icon"><i class="flaticon-h264"></i></a><a href="/torrent/5138248/SpiderMan-No-Way-Home-2021-HD-TS-x264-AAC/">SpiderMan No Way Home 2021 HD-TS x264 AAC</a></td>
<td class="coll-2 seeds">262</td>
<td class="coll-3 leeches">41</td>
<td class="coll-date">Feb. 3rd '22</td>
<td class="coll-4 size mob-vip">1,002.1 MB<span class="seeds">262</span></td>
<td class="coll-5 vip"><a href="/user/mazemaze16/">mazemaze16</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/4/0/" class="icon"><i class="flaticon-video-dual-sound"></i></a><a href="/torrent/2554075/SpiderMan-Homecoming-2017-KiSS-BluRay-720p-HD-AAC-Hindi-Eng-French-Subs-mkv/">SpiderMan Homecoming 2017 [KiSS] BluRay 720p HD AAC [Hindi-Eng-French] Subs.mkv</a><span class="comments"><i class="flaticon-message"></i>5</span></td>
<td class="coll-2 seeds">257</td>
<td class="coll-3 leeches">19</td>
<td class="coll-date">Oct. 30th '17</td>
<td class="coll-4 size mob-uploader">2.5 GB<span class="seeds">257</span></td>
<td class="coll-5 uploader"><a href="/user/KissToonTv/">KissToonTv</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/42/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/416176/The-Amazing-SpiderMan-2012-720p-BrRip-x264-YIFY/">The Amazing SpiderMan (2012) 720p BrRip x264 - YIFY</a><span class="comments"><i class="flaticon-message"></i>1</span></td>
<td class="coll-2 seeds">236</td>
<td class="coll-3 leeches">27</td>
<td class="coll-date">Oct. 25th '12</td>
<td class="coll-4 size mob-vip">899.3 MB<span class="seeds">236</span></td>
<td class="coll-5 vip"><a href="/user/YIFY/">YIFY</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/2/0/" class="icon"><i class="flaticon-divx"></i></a><a href="/torrent/5177531/SpiderMan-No-Way-Home-2021-HD-TS-x264-AAC/">SpiderMan.No.Way.Home.2021.HD-TS.x264.AAC</a><span class="comments"><i class="flaticon-message"></i>4</span></td>
<td class="coll-2 seeds">223</td>
<td class="coll-3 leeches">21</td>
<td class="coll-date">Mar. 10th '22</td>
<td class="coll-4 size mob-vip">1,001.1 MB<span class="seeds">223</span></td>
<td class="coll-5 vip"><a href="/user/mazemaze16/">mazemaze16</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/70/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/5184694/Spider-Man-No-Way-Home-2021-2160p-BluRay-SDR-Eng-Hin-Tel-Tam-10bit-HQ-HEVC-By-DCR/">Spider-Man.No.Way.Home.2021.2160p.BluRay.SDR.[Eng-Hin-Tel-Tam].10bit.HQ.HEVC.By.DCR</a><span class="comments"><i class="flaticon-message"></i>9</span></td>
<td class="coll-2 seeds">196</td>
<td class="coll-3 leeches">57</td>
<td class="coll-date">Mar. 16th '22</td>
<td class="coll-4 size mob-trial-uploader">8.4 GB<span class="seeds">196</span></td>
<td class="coll-5 trial-uploader"><a href="/user/DCREncodes/">DCREncodes</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/42/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/5180737/Spider-Man-No-Way-Home-2021-720p-BluRay-x264-NeZu/">Spider-Man.No.Way.Home.2021.720p.BluRay.x264-NeZu</a><span class="comments"><i class="flaticon-message"></i>2</span></td>
<td class="coll-2 seeds">148</td>
<td class="coll-3 leeches">50</td>
<td class="coll-date">Mar. 13th '22</td>
<td class="coll-4 size mob-vip">1.5 GB<span class="seeds">148</span></td>
<td class="coll-5 vip"><a href="/user/NeZu/">NeZu</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/42/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/5179515/Spider-Man-No-Way-Home-2021-Dual-Audio-Hindi-Cleaned-1080p-BRRip-x264-ProLover/">Spider-Man No Way Home (2021) Dual Audio Hindi (Cleaned) 1080p BRRip x264 - ProLover</a><span class="comments"><i class="flaticon-message"></i>2</span></td>
<td class="coll-2 seeds">143</td>
<td class="coll-3 leeches">68</td>
<td class="coll-date">Mar. 12th '22</td>
<td class="coll-4 size mob-uploader">1.6 GB<span class="seeds">143</span></td>
<td class="coll-5 uploader"><a href="/user/prolover/">prolover</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/4/0/" class="icon"><i class="flaticon-video-dual-sound"></i></a><a href="/torrent/2560074/SpiderMan-Homecoming-2017-KiSS-BluRay-720p-Multi-Subs-Hindi-Tamil-French-Eng/">SpiderMan Homecoming 2017 [KiSS] BluRay.720p.Multi-Subs.[Hindi.Tamil.French-Eng]</a><span class="comments"><i class="flaticon-message"></i>18</span></td>
<td class="coll-2 seeds">139</td>
<td class="coll-3 leeches">30</td>
<td class="coll-date">Nov. 2nd '17</td>
<td class="coll-4 size mob-uploader">1.2 GB<span class="seeds">139</span></td>
<td class="coll-5 uploader"><a href="/user/KissToonTv/">KissToonTv</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/42/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/3583638/Marvel-Animated-Movies-Collection/">Marvel Animated Movies Collection</a><span class="comments"><i class="flaticon-message"></i>16</span></td>
<td class="coll-2 seeds">136</td>
<td class="coll-3 leeches">235</td>
<td class="coll-date">Feb. 10th '19</td>
<td class="coll-4 size mob-uploader">24.4 GB<span class="seeds">136</span></td>
<td class="coll-5 uploader"><a href="/user/ShinkFive/">ShinkFive</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/54/0/" class="icon"><i class="flaticon-h264"></i></a><a href="/torrent/3868632/Spider-Man-Far-From-Home-2019-New-Source-Hindi-720p-CAMRip-MP3-x264-HDWebMovies/">Spider-Man Far From Home (2019) [New Source] Hindi 720p CAMRip MP3 x264 [HDWebMovies]</a><span class="comments"><i class="flaticon-message"></i>4</span></td>
<td class="coll-2 seeds">129</td>
<td class="coll-3 leeches">25</td>
<td class="coll-date">Jul. 6th '19</td>
 <td class="coll-4 size mob-uploader">1.0 GB<span class="seeds">129</span></td>
<td class="coll-5 uploader"><a href="/user/Suryadipta/">Suryadipta</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/70/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/5185361/Spider-Man-No-Way-Home-2022-1080p-Bluray-Multi-Audio-10bit-DTS-HD-MA-5-1-x265-HashMiner/">Spider-Man.No.Way.Home.2022.1080p.Bluray.Multi-Audio.10bit.DTS-HD.MA.5.1.x265.[HashMiner]</a><span class="comments"><i class="flaticon-message"></i>2</span></td>
<td class="coll-2 seeds">120</td>
<td class="coll-3 leeches">43</td>
<td class="coll-date">Mar. 17th '22</td>
<td class="coll-4 size mob-uploader">9.0 GB<span class="seeds">120</span></td>
<td class="coll-5 uploader"><a href="/user/HashMiner/">HashMiner</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/42/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/5190272/SpiderMan-No-Way-Home-2021-1080P-H264-Ita-Eng-Ac3-5-1-Sub-Ita-Eng-SnakeSPL-MIRCrew/">SpiderMan: No Way Home (2021) 1080P H264 Ita Eng Ac3 5.1 Sub Ita Eng SnakeSPL MIRCrew</a></td>
<td class="coll-2 seeds">116</td>
<td class="coll-3 leeches">8</td>
<td class="coll-date">Mar. 21st '22</td>
<td class="coll-4 size mob-uploader">3.9 GB<span class="seeds">116</span></td>
<td class="coll-5 uploader"><a href="/user/SnakeSPL1337/">SnakeSPL1337</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/42/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/3644714/Spider-Man-Into-the-Spider-Verse-2018-720p-BluRay-x264-NeZu/">Spider-Man.Into.the.Spider-Verse.2018.720p.BluRay.x264-NeZu</a></td>
<td class="coll-2 seeds">115</td>
<td class="coll-3 leeches">37</td>
<td class="coll-date">Mar. 12th '19</td>
<td class="coll-4 size mob-vip">1.3 GB<span class="seeds">115</span></td>
<td class="coll-5 vip"><a href="/user/NeZu/">NeZu</a></td>
</tr>
<tr>
<td class="coll-1 name"><a href="/sub/42/0/" class="icon"><i class="flaticon-hd"></i></a><a href="/torrent/416521/The-Amazing-SpiderMan-2012-1080p-BrRip-x264-YIFY/">The Amazing SpiderMan (2012) 1080p BrRip x264 - YIFY</a><span class="comments"><i class="flaticon-message"></i>4</span></td>
<td class="coll-2 seeds">97</td>
<td class="coll-3 leeches">33</td>
<td class="coll-date">Oct. 26th '12</td>
<td class="coll-4 size mob-vip">2.0 GB<span class="seeds">97</span></td>
<td class="coll-5 vip"><a href="/user/YIFY/">YIFY</a></td>
</tr>
</tbody>
</table>
</div>
<div class="pagination">
<ul>
<li class="active"><a href="/search/spiderman/1/">1</a></li>
<li><a href="/search/spiderman/2/">2</a></li>
<li><a href="/search/spiderman/3/">3</a></li>
<li><a href="/search/spiderman/4/">4</a></li>
<li><a href="/search/spiderman/5/">5</a></li>
<li><a href="/search/spiderman/6/">6</a></li>
<li><a href="/search/spiderman/7/">7</a></li>
<li><a href="/search/spiderman/2/">&gt;&gt;</a></li>
<li class="last"><a href="/search/spiderman/23/">Last</a></li>
</ul>
</div>
</div>
<br>
<center>
<div class="l190adf4ec5fde22551dc01b20d52b4030ea333b6">
<a href="/anoydl9-VpnHideMe"><img src="/css/images/axbotx2.png" /></a>
</div>
</center>
<br> </div>
</div>
</div>
</main>
<footer>
<div class="bitcoin">
<div class="bitcoin-icon-wrap">
<span class="bitcoin-icon"><i class="flaticon-bitcoin red"></i></span>
</div>
<span class="bitcoin-text"><span>Bitcoin Donate: </span><a href="bitcoin:3Q1337xL2i6jXrXqZ5aMfhN4wp366GQc44">3Q1337xL2i6jXrXqZ5aMfhN4wp366GQc44</a></span>
</div>
<a class="scroll-top" href="#"><i class="flaticon-up"></i></a>
<ul>
<li><a href="/">Home</a></li>
<li class="active"><a href="/home/">Full Home Page</a></li>
<li><a href="/contact">Dmca</a></li>
<li><a href="/contact">Contact</a></li>
</ul>
<p class="info">1337x 2007 - 2022</p>
</footer>
<script src="/js/jquery-1.11.0.min.js"></script>
<script src="/js/jquery-ui.js"></script>
<script src="/js/auto-searchv2.js"></script>
<script src="/js/fancySelect.js"></script>
<script src="/js/main.js"></script>
<script>

  $(window).load(function() {
    $('.select').fancySelect();
  });
</script>
</body>
</html>
	`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, row *goquery.Selection) {
		if attr, exists := row.Attr("href"); exists {
			// headings = append(headings, attr)
			fmt.Println(attr)
		}

	})

	// doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
	// 	tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
	// 		rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
	// 			headings = append(headings, tableheading.Text())
	// 		})
	// 		rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
	// 			row = append(row, tablecell.Text())
	// 		})
	// 		rows = append(rows, row)
	// 		row = nil
	// 	})
	// })
	// fmt.Println("####### headings = ", len(headings), headings)
	// fmt.Println("####### rows = ", len(rows), rows)
}

func main() {

	url := "https://1337x.to/search/spiderman/1/"

	// getinfo(url)
	links := getinfo(url)

	// print out the results
	for _, link := range links {
		fmt.Println(link)
	}
	// random(url)

	// learnJquery()
}
