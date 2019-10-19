package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)        // --> & gives you the address. Like 0xc4200a4000: address of memory
	fmt.Printf("%T\n", a)  // --> int
	fmt.Printf("%T\n", &a) // --> *int

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an adress when you have the adress
	fmt.Println(*&a)

	*b = 43
	fmt.Println(*b)
	fmt.Println(a)

	fmt.Println("========== When use pointer =========")
	x := 0
	foo(x)
	fmt.Println("x before: ", x) // --> 0
	fmt.Println("&x before: ", &x)
	bar(&x)
	fmt.Println("x after passing address", &x)
	fmt.Println("x after passing address", x)

}

func foo(y int) { // parameter y is passed by value
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func bar(y *int) { // parameter y is passed by value
	fmt.Println("y before: ", y)
	fmt.Println("*y before: ", *y)
	*y = 43
	fmt.Println("y after: ", y)
	fmt.Println("*y after: ", *y)
}
