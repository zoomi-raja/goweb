package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	itration int
)

// SitemapIndex will be exported
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

//News is to parse
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

func getNews(url string, length int, c chan News) {
	var n News
	resp, err := http.Get(strings.TrimSpace(url))
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	itration++
	fmt.Println(itration)
	c <- n
	if itration == length {
		close(c)
	}
}

func parse() map[string]NewsMap {
	var s SitemapIndex
	newsMap := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	/*stringBody := string(bytes)
	fmt.Println(stringBody) */
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)
	chanelLength := len(s.Locations)
	c := make(chan News, chanelLength)
	for i := 0; i < chanelLength; i++ {
		go getNews(s.Locations[i], chanelLength, c)
	}
	for ch := range c {
		for idx, _ := range ch.Keywords {
			newsMap[ch.Titles[idx]] = NewsMap{ch.Keywords[idx], ch.Locations[idx]}
		}
	}
	return newsMap
}

func testHowReaderWork() {
	r := strings.NewReader("this is string to be read by reader of golang")
	if s, err := ioutil.ReadAll(r); err != nil {
		fmt.Println("issue in reading")
	} else {
		fmt.Fprintf(os.Stderr, "%s", s) //os.Stdout
	}
}
