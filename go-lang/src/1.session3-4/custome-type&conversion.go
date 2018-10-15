package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Impossible as different type
	// a = b

	// Conversion
	// int(b) will convert hotdog b to int
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
