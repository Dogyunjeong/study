package main

import (
	"fmt"
)

func main() {
	var y string

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"

	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
