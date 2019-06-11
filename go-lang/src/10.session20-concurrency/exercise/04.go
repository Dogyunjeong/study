package main

import (
	"fmt"
	"sync"
)

func main() {
	// Fix the race condition you created in the previous exercise by using a mutex
	// 	it makes sense to remove runtime.Gosched()
	var wg sync.WaitGroup
	var mu sync.Mutex
	increment := 0
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := increment
			v++
			increment = v
			fmt.Println(increment)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value: ", increment)
}
