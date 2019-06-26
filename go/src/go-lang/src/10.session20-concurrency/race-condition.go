package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	counter := 0

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)  // Sleep go routine for certain time
			runtime.Gosched()
			v++
			counter = v
			wg.Done() // Go routine finish
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait() // wait unitle finish all go routine

	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count: ", counter)
}

// run on bash `go run --race thisfile.go`
