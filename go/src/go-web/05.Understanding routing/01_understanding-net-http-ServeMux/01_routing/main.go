package main

import (
	"io"
	"net/http"
)

type hotdog int

// It makes hotdog as a implicitly implemented handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Routing. There is better way to do it. This is just for simple example
	// Not elegant -> bit elegant -> very elegant ...
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	// d has two underlying type. hotdog and handler
	var d hotdog
	http.ListenAndServe(":8080", d)
}
