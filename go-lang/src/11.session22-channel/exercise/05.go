package main

import (
	"fmt"
)

/**
Show the comma ok idiom starting with
func main() {
	c := make(chan int)

	v, ok :=
		fmt.Println(v, ok)

	close(c)

	v, ok =
		fmt.Println(v, ok)
}
*/
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
