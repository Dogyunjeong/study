package main

import (
	"fmt"
)

func main() {
	// Exercise 1
	fmt.Println("--------Exercise 1---------")

	ex := 13
	fmt.Printf("%b\n", ex)
	fmt.Printf("%d\n", ex)
	fmt.Printf("%x\n", ex)

	// Exercise 2
	// Use comparison operator and assign their values
	fmt.Println("--------Exercise 2---------")

	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println(a, b, c, d, e, f)

	// Exercise 3
	// Create and print untyped constant and typed constant
	fmt.Println("--------Exercise 3---------")

	const (
		testUntyped     = 42 // untyped
		testTyped   int = 43 // typed
	)
	fmt.Println(testUntyped, "  ", testTyped) // ----> 42  43

	// Exercise 4
	// Assign int and prints in decimal, binary and hex. Reapeat print after shifting bits of that int over 1
	// We need bit shifting where the systems that do not have floating point support for nearly all arithmetic. For polynomial arithmetic, generates hash and etc.
	fmt.Println("--------Exercise 4---------")
	testInt := 42
	fmt.Printf("%b\n", testInt)
	fmt.Printf("%d\n", testInt)
	fmt.Printf("%x\n", testInt)

	bitShifted := testInt << 1

	fmt.Printf("%b\n", bitShifted)
	fmt.Printf("%d\n", bitShifted)
	fmt.Printf("%x\n", bitShifted)

	// Exercise 5
	// Create a variable of type string using a raw string literal. Print it.
	// WORD-raw string literal: https://en.wikipedia.org/wiki/String_literal#Raw_strings
	fmt.Println("--------Exercise 5---------")

	rawStr := `test raw string /hi \hello`
	fmt.Println(rawStr)

	// Exercise 6
	// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
	fmt.Println("--------Exercise 6---------")
	const (
		next1 = 2017 + iota
		next2 = 2017 + iota
		next3 = 2017 + iota
		next4 = 2017 + iota
	)

	fmt.Println(next1, next2, next3, next4)

	// Exercise 7
	// test: https://docs.google.com/forms/d/e/1FAIpQLSfjhxXjo0r_OsVys58B1lVs35CLPpneVcjiEKTPsLuQs4mftA/viewform
}
