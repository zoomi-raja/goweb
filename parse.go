package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func parse() {
	var s SitemapIndex
	var n News
	newsMap := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	/*stringBody := string(bytes)
	fmt.Println(stringBody) */
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)
	for _, url := range s.Locations {
		resp, err := http.Get(strings.TrimSpace(url))
		if err != nil {
			fmt.Println(err)
		}
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx := range n.Keywords {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	fmt.Println(newsMap)
}

func testHowReaderWork() {
	r := strings.NewReader("this is string to be read by reader of golang")
	if s, err := ioutil.ReadAll(r); err != nil {
		fmt.Println("issue in reading")
	} else {
		fmt.Fprintf(os.Stderr, "%s", s)
	}
}
