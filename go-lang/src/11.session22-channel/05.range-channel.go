package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	for v := range c { // waiting until channel is close when there is range loop for channel
		fmt.Println(v)
	}
	// channel pull will block code until channel gives value. so even there is no waitgroup it will works

	fmt.Println("about to exit")
}
