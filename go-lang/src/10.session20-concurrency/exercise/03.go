package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Using goroutines, create an incrementer program
	// 	have a variable to hold the incrementer value
	// 	launch a bunch of goroutines
	// 		each goroutine should
	// 		read the incrementer value
	// 		store it in a new variable
	// 		yield the processor with runtime.Gosched()
	// 		increment the new variable
	// 		write the value in the new variable back to the incrementer variable
	// use waitgroups to wait for all of your goroutines to finish
	// the above will create a race condition.
	// Prove that it is a race condition by using the -race flag
	// if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
	var wg sync.WaitGroup

	increment := 0
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := increment
			runtime.Gosched()
			v++
			increment = v
			fmt.Println(increment)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value: ", increment)
}
