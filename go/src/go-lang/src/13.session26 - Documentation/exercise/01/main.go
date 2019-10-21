package main

import (
	"fmt"
	"lib/dog"
)

type canine struct {
	age  int
	name string
}

func (canine *canine) getHumanAge() int {
	return dog.Years(canine.age)
}

func main() {
	myDog := canine{
		age:  7,
		name: "Puppy",
	}
	fmt.Println(myDog.getHumanAge())
}
