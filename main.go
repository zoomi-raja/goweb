package main

import (
	"fmt"
	"net/http"
)

type myHandler string

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("peep zoomi its working"))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("peep zoomi its working user handler func"))
}
func main() {
	// var handler myHandler
	// http.Handle("/", handler)
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
