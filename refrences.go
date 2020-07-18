package main

import (
	"fmt"
	"net/http"
)

type myHandler string

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("direct zoomi its working with method " + r.Method))
	// or becuase inside Fprintf its using w.Write
	fmt.Fprintf(w, `<h1>fmt zoomi its working with method %s<h1>
	<span>peeeep</span>`, r.Method)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("peep zoomi its working user handler func"))
}
func refrence() {
	var handler myHandler
	http.Handle("/handler", handler) //http.handlerfunc(handlerFunc)
	http.HandleFunc("/", handlerFunc)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("error accured", err)
	}
}
