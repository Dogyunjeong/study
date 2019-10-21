package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type hotdog int

func main() {
	foo() // have to specify parameters even there is no parameters. ()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("========= Variadic parameter ==========")
	x2 := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is", x2)

	fmt.Println("========= Unfurling a slice ==========")
	xs := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x4 := sum(xs...)
	fmt.Println("The total is", x4)

	fmt.Println("========= defer ==========")
	defer foo() // run this code after all line is done in this block
	bar("Bar")

	fmt.Println("========= methods ==========")
	sa1 := secretAgent{ // type of sa1 are secretAgent and human
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println("========= interface & polymorphism ==========")
	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	fmt.Println(p1)
	bar2(sa1) // sa1 is type secretAgent
	bar2(sa2)
	bar2(p1) // p is type person

	// conversion
	var x5 hotdog = 42
	fmt.Println(x5)
	fmt.Printf("%T\n", x5)
	var y5 int
	y5 = int(x5)
	fmt.Println(y5)
	fmt.Printf("%T\n", y5)
	// assertion
}

//  func (r receiver) identifier(parameters) (returns(s)) { ... }

func foo() { // no receiver, parameters and return
	fmt.Println("hello from foo")
}

// Everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo,", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := true
	return a, b
}

func sum(x ...int) int { // x is slice of int
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, "we are now adding,", v, "to the total which is now", sum)
	}
	fmt.Println("The total is,", sum)
	return sum
}

// func (r receiver) identifier(parameters) (returns(s)) {...}
// when there is receiver, the function will be attached to receiver

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// interface allow more than one type
// polymorphism
type human interface {
	speak()
}

func bar2(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrr ", h.(person).first) // assetion
	case secretAgent:
		fmt.Println("I was passed into barrrrr ", h.(secretAgent).first)

	}
	fmt.Println("I was passed into bar ", h)

}
