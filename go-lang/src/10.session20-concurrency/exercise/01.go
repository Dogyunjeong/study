package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// in addition to the main goroutine, launch two additional goroutines
	// each additional goroutine should print something out
	// use waitgroups to make sure each goroutine finishes before your program exists

	fmt.Println("begin CPUs: ", runtime.NumCPU())
	fmt.Println("begin Goroutines: ", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		runtime.Gosched()
		fmt.Println("Hello routine 1")
		wg.Done()
	}()

	go func() {
		runtime.Gosched()
		fmt.Println("Hello routine 2")
		wg.Done()
	}()

	fmt.Println("mid CPUs: ", runtime.NumCPU())
	fmt.Println("mid Goroutines: ", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("about exit")
	fmt.Println("end CPUs: ", runtime.NumCPU())
	fmt.Println("end Goroutines: ", runtime.NumGoroutine())
}
