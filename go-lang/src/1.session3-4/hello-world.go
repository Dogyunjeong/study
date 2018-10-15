package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello everyone, this is the most popular course")

	foo()

	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()

	declare()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}

func declare() {
	x := 42 // short declaralation
	fmt.Println(x)
	x = 99 // reassign
	fmt.Println(x)
	y := 100 + 4 // statement with expression
	fmt.Println(y_)
}
