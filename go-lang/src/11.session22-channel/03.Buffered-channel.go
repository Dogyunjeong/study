package main

import "fmt"

// It will not run
func main() {
	c := make(chan int, 1) // will run with buffered channer

	c <- 42 // Chanel block the code
	// c <- 43 // will block even buffered channel to pull one more time
	fmt.Println(<-c)
	fmt.Println("end")

}
