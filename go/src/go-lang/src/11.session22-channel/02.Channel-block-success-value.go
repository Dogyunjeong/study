package main

import "fmt"

func main() {
	c := make(chan int, 1) // burferred channel

	// Channel can run
	go func() {
		c <- 42 // Block here
	}()

	fmt.Println(<-c)

}
