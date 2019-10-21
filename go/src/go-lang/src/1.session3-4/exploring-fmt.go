package main

import (
	"fmt"
)

var y = 42

func main() {
	// refer to format package in golang.org
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	// go lang escape represent \n newline \t tab
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	// String printing is assignable
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)

	// print default value
	fmt.Printf("%v\n", y)
}
