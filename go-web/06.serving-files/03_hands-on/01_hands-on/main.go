package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

type foo int

func (f foo) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo bar")
}

type bar int

func (b bar) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

type chien int

func (c chien) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}

func main() {
	var f foo
	var b bar
	var c chien
	mux := http.NewServeMux()
	mux.Handle("/foo", f)
	mux.Handle("/bar", b)
	mux.Handle("/dog.jpg", c)
	http.ListenAndServe(":8080", mux)
}
