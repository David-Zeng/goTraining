package main

import (
	"fmt"
	"github.com/David-Zeng/goTraining/myFunctions"
)

type person struct {
	first string
	last string
	age int
}
func (p person) speak() {
	fmt.Println("I am", p.first, ". I am", p.age, "years old.")
}


func main() {

	p1 := person{
		first: "James",
		last: "Bonds",
		age: 38,
	}
	myFunctions.ObjectDescribe(p1)

	p1.speak()

}
