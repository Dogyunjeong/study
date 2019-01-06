/**
lunches 10 goroutines
	each goroutine adds 10 numbers to a channel
pull the numbers off the channel and print them
*/

package main

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
