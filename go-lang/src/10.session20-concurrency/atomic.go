package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	var counter int64 // package atomic needs int64

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched() // switch between go routine
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done() // Go routine finish
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait() // wait until finish all go routine

	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count: ", counter)
}

// run on bash `go run thisfile.go`
