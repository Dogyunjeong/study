package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat
	// It is not still elegant way
	mux := http.NewServeMux()
	mux.Handle("/dog/", d) // this will catch /dog/* as well because of fore slash
	mux.Handle("/cat", c)  // this will only catch /cat  not cat/something

	http.ListenAndServe(":8080", mux)
}

/**
*  Usnig defaultServeMux with passing nil
*  Then we will use Handle and HandleFunc
**/
// HandleFunc has two underlying type. HandlerFunc and Handler
func e(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func f(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}
func forStudying() {
	// Attached to default serveMux
	http.HandleFunc("/dog/", e)
	http.HandleFunc("/cat", f)
	// HandlerFunc is just helper function to create handler
	http.Handle("/cat", http.HandlerFunc(f))

	http.ListenAndServe(":8080", nil) // It will use default serveMux  https://godoc.org/net/http#ListenAndServe
}
