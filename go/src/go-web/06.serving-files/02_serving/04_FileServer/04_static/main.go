package main

import (
	"log"
	"net/http"
)

func main() {
	// http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
