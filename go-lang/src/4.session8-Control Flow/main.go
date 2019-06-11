package main

import (
	"fmt"
)

func main() {
	//  Go doesn't have 'while' loop
	/*
	*	for initial statement; condition; post {
	*
	*  }
	 */
	//  nesting loop
	//  outer loop
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		// inner loop
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}

	// How do like while statement
	fmt.Println("====== How do like while ======")
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

	// for with if statement
	fmt.Println("====== for with if stmt ======")
	y := 1
	for {
		y++
		if y > 100 {
			break
		}
		// how about when y is 0. eg  0 / 2, 0%2
		if y%2 != 0 {
			continue
		}

		fmt.Println(y)
	}

	fmt.Println("====== remainder - modulo operation ======")
	a := 43 / 40
	fmt.Println(a)
	b := 43 % 40
	fmt.Println(b)

	fmt.Println("====== printing ascii ======")
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i) // number  hexadecimal  ascii  pointing-character
	}

	fmt.Println("====== if statement ======")
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true { // true is pre declaration constant and ! is not operator
		fmt.Println("003")
	}
	if !false { // double negative
		fmt.Println("004")
	}
	if 2 == 2 {
		fmt.Println("005")
	}
	if 2 != 2 {
		fmt.Println("006")
	}
	if !(2 == 2) {
		fmt.Println("007")
	}
	if !(2 != 2) {
		fmt.Println("008")
	}

	fmt.Println("====== initialization statement ======")
	// initX is valid inside of if statement as declaring in scope
	if initX := 42; initX == 40 { // We don't need to put ;. When we want to write tow statement we should add ;
		fmt.Println("test inside of if stmt")
	}

	fmt.Println("====== else if and else statement ======")
	elseX := 42
	// initX is valid inside of if statement as declaring in scope
	if elseX == 40 { // We don't need to put ;. When we want to write tow statement we should add ;
		fmt.Println("our value was 40")
	} else if elseX == 41 {
		fmt.Println("our value was not 41")
	} else {
		fmt.Println("our value was not 40, 41")
	}

	fmt.Println("====== repeat loop, modulo, conditional ======")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("====== switch statement ======")
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	// it will be not reached without fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	case (7 == 9): // will be ran due to fallthrough
		fmt.Println("not true1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
		fallthrough
	default: // will be ran due to fallthrough
		fmt.Println("this is default")
	}
	fmt.Println("====== switch statement 2 ======")
	switch "Bond" {
	case "Moneypenny", "Bond", "Do No": // multiple condition
		fmt.Println("miss money or bond or do no")
	case "M":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is q")
	}

	switch { // missing switch statement is equivalent to boolean true
	case false:
		fmt.Println("Should not print")
	case true:
		fmt.Println("Should print")
	}

	fmt.Println("====== conditional logic operating ======")
	fmt.Printf("true && true\t %v\n", true && true)
	fmt.Printf("true && false\t %v\n", true && false)
	fmt.Printf("true || true\t %v\n", true || true)
	fmt.Printf("true || false\t %v\n", true || false)
	fmt.Printf("!true\t\t %v\n", !true)
	fmt.Printf("!false\t\t %v\n", !false)
}
