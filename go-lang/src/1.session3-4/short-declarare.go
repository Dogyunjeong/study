package main

import (
	"fmt"
)

// DECLARE the variable "p"
// ASSIGN the value 43
// declare & assign = initialization
var p = 43 // p is package level

// DECLARE tehre is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABlE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
var z int

func main() {
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 77
	fmt.Println(y)
	i := "Bond, James"
	fmt.Println(i)

	var h = 72
	fmt.Println(h)
}
