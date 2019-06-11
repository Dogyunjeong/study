package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)

	fmt.Println("======= embede struct =======")
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)

	fmt.Println("======= assign value =======")
	var x, y int
	x = 42
	y = 43
	fmt.Println(x, y)
	x, y = 44, 45
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

	fmt.Println("======= anonymous struct =======")
	p3 := person{ // composite literal: {}
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p3)

	// anonymous struct
	// As it doesn't have name
	p4 := struct {
		first string
		last  string
		age   int
	}{ // composite literal: {}
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p4)

	fmt.Println("======= housekeeping OOP =======")
	// we create VALUES of a certain TYPE that are stored in VARIABLES
	// and those VARIABLES have IDENTIFIERS
	var x int // KEYWORD IDENTIFIER UNDERLYING_TYPE
	p5 := person{
		first: "James",
		last:  "Bond",
	}

}
