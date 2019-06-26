package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// Fix the race condition you created in exercise #4 by using package atomic
	//	code: https://github.com/GoesToEleven/go-programming
	var wg sync.WaitGroup
	var increment int32
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt32(&increment, 1)
			fmt.Println(atomic.LoadInt32(&increment))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value: ", increment)
}
