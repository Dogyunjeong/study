package main

import (
	"fmt"
)

// https://golang.org/ref/spec#Types
// Create your own type. Have the underlying type be an int.
// create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
// in func main
// print out the value of the variable “x”
// print out the type of the variable “x”
// assign 42 to the VARIABLE “x” using the “=” OPERATOR
// print out the value of the variable “x”

// underlying type mine will be type int
// type mine = int

type mine int

// compiler will assign zero value. we don't assign yet. just declare x with type mine
var x mine

func main() {
	fmt.Printf("%v\n", x)
	// 0
	fmt.Printf("%T\n", x)
	// main.mine
	x = 42
	fmt.Printf("%v\n", x)
}
