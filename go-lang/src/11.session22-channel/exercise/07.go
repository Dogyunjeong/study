/**
lunches 10 goroutines
	each goroutine adds 10 numbers to a channel
pull the numbers off the channel and print them
*/
/**
My failed solution
func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c)

	for k := range c {
		fmt.Println(k)
	}
}
*/

package main

import "fmt"

func main() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}
