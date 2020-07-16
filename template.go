package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type newsAggPage struct {
	Title string
	News  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hurray still works</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	p := newsAggPage{"Amazing News Aggregator", "some news"}
	if t, err := template.ParseFiles(cwd + "/src/gotut/templates/basictemplating.html"); err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, p)
	}
}

func templateFunc() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error :", err)
	}
}
