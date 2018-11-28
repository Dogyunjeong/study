package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("=======Bool=======")

	a := 7
	b := 42
	fmt.Println(a == b)
	fmt.Println(a != b)

	// Number
	fmt.Println("=======Number=======")

	x := 42
	y := 42.234234
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	// String
	// String type represents the set of string values. Sequence of bytes
	fmt.Println("=======String=======")

	s := `"Hello, Play ground`
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// Slice string(it is sequence of bytes). each number is code point in UTF-8 code scheme
	bs := []byte(s) // --> [34 72 101 108 108 111 44 32 80 108 97 121 32 103 114 111 117 110 100]
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		// print UTF8 code point
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("  -- with range -- ")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v) // --> such like "at index position 9 we have hex 0x6c"
	}

	/*
		numeral system
		binary, hexadecimal, and etc
	*/

	fmt.Println("=======Const=======")
	const i = 42            // untyped constant
	const j float64 = 42.78 // typed constant
	const k = "James Bond"

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", j)
	fmt.Printf("%T\n", k)

	// pre declared identifier iota
	// It will increase automatically
	fmt.Println("=======iota=======")

	const (
		t = iota
		g
		h
	)

	const (
		d = iota
		e
		f
	)
	fmt.Println(t)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", t)
	fmt.Printf("%T\n", g)
	fmt.Printf("%T\n", h)

	fmt.Println("=======bit shifting=======")

	bitShifting := 2
	shifted := bitShifting << 1
	fmt.Printf("%d\t\t%b\n", bitShifting, bitShifting) // --> 100
	fmt.Printf("%d\t\t%b\n", shifted, shifted)         // --> 1000

	kb := 1024
	mb := kb * 1024

	fmt.Println(mb)

}
