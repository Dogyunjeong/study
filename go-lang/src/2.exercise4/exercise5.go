package main

import (
	"fmt"
)

// Building on the code from the previous example
// at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
// eg:
// 		type hotdog int
// 		var x hotdog
// 		var y int

// in func main
// this should already be done
// print out the value of the variable “x”
// print out the type of the variable “x”
// assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
// print out the value of the variable “x”
// now do this
// now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
// then use the “=” operator to ASSIGN that value to “y”
// print out the value stored in “y”
// print out the type of “y”

type mine int

// compiler will assign zero value. we don't assign yet. just declare x with type mine
var x mine
var y int

func main() {
	fmt.Printf("%v\n", x) // 0
	fmt.Printf("%T\n", x) // main.mine
	x = 42
	fmt.Printf("%v\n", x)
	y = int(x)
	fmt.Println(x)
	fmt.Printf("%T\n", y)
}
