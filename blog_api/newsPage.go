package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type newsAgg struct {
	Title string
	News  map[string]NewsMap
}

func newsPageHandler(w http.ResponseWriter, r *http.Request) {
	news := parse()

	if t, err := template.ParseFiles("./templates/newsTemplate.html"); err != nil {
		fmt.Println("error", err)
	} else {
		p := newsAgg{"Amazing News Aggregator", news}
		t.Execute(w, p)
	}
}
func newsPage() {

	http.HandleFunc("/", newsPageHandler)
	http.ListenAndServe(":3000", nil)
}
