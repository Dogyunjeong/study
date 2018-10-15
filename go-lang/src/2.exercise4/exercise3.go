package main

import (
	"fmt"
)

// At the package level scope, assign the following values to the three variables
// for x assign 42
// for y assign “James Bond”
// for z assign true
// in func main
// use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
// print out the value stored by variable “s”

// A: Zero value

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\n", x, y, z)

	fmt.Println(s)
}
