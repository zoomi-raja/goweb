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
	Locations []Location `xml:"sitemap"`
}

//Location peep peep
type Location struct {
	Loc string `xml:"loc"`
}

func (L Location) String() string {
	return fmt.Sprintf(L.Loc)
}

func parse() {
	testHowReaderWork()
	peep, k := true, 0
	for peep {
		k++
		if k > 5 {
			peep = false
		}
		fmt.Println("ji ")
	}

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	/*stringBody := string(bytes)
	fmt.Println(stringBody) */
	resp.Body.Close()
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	for _, url := range s.Locations {
		fmt.Println(url)
	}
}

func testHowReaderWork() {
	r := strings.NewReader("this is string to be read by reader of golang")
	if s, err := ioutil.ReadAll(r); err != nil {
		fmt.Println("issue in reading")
	} else {
		fmt.Fprintf(os.Stderr, "%s", s)
	}
}
