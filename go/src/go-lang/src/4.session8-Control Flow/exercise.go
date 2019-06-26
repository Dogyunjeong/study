package main

import (
	"fmt"
)

func main() {
	fmt.Println("===== exercise 1 ======")
	// exercise #1
	// Print every number form 1 to 1000
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}

	fmt.Println("===== exercise 2 ======")
	// exercise #2
	// Print every rune code point of the uppercase alphabet three times
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n ", i)
		}
	}

	fmt.Println("===== exercise 3 ======")
	// exercise #3
	// create for loop using for condition {}
	bd := 1985
	for bd <= 2017 {
		fmt.Println(bd)
		bd++
	}

	fmt.Println("===== exercise 4 ======")
	// exercise #4
	// create for loop using for {}
	bd2 := 1985
	for {
		if bd2 > 2017 {
			break
		}
		fmt.Println(bd2)
		bd2++
	}

	fmt.Println("===== exercise 5 ======")
	// exercise #5
	//  Print out the remainder (modulus) whichi is found for each number between 10 and 100 when it is divided by 4
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divide by 4, the remiander (aka modulus) is %v\n", i, i%4)
	}

	fmt.Println("===== exercise 6 ======")
	// exercise #6
	// create a program that shows the "if statement" in action
	x := "James Bond"
	if x == "James Bond" {
		fmt.Println(x)
	}

	fmt.Println("===== exercise 7 ======")
	// exercise #7
	// Building on the previous hands-on exercise, create a program that uses "else if" and "else"
	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("Else if: ", x)
	} else {
		fmt.Println("Else")
	}

	fmt.Println("===== exercise 8 ======")
	// exercise #8
	// create a program with switch statement with no switch expression specified
	switch {
	case false:
		fmt.Println("Should not print")
	case true:
		fmt.Println("Should print")
	}

	fmt.Println("===== exercise 9 ======")
	// exercise #9
	// create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER "favSport"
	favSport := "hiking"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "hiking":
		fmt.Println("go to canada!")
	}

	fmt.Println("===== exercise 10 ======")
	// exercise #10
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
