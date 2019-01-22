package main

import (
	"fmt"
	"log"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}
	foo(c1)
}

func foo(e error) {
	log.Println("foo ran - ", e, "\n", e.(customErr).info) // to access info we need assertion as e is type error. error has not `info` property
	log.Println("foo ran - ", e, "\n")
}
