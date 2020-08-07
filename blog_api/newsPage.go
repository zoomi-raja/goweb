package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type newsAgg struct {
	Title string
	News  map[string]NewsMap
}

func newsPageHandler(w http.ResponseWriter, r *http.Request) {
	news := parse()
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if t, err := template.ParseFiles(cwd + "/src/gotut/templates/newsTemplate.html"); err != nil {
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
