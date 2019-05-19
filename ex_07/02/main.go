package main

import (
	"fmt"
	"github.com/David-Zeng/goTraining/myFunctions"
)

type person struct {
	first string
	last string
	address string
}

func (p *person) changeMe(updatedAddress string) {
	p.address = updatedAddress
}

func main() {
	fmt.Println("pointer ex 02")

	p1 := person{
		first: "Joanne",
		last: "Jason",
		address: "Sydney, AU",
	}

	myFunctions.ObjectDescribe(p1)

	p1.changeMe("Melbourn")

	myFunctions.ObjectDescribe(p1)

	(&p1).first = "Bob"
	myFunctions.ObjectDescribe(p1)

	p1.first = "Jason"
	myFunctions.ObjectDescribe(p1)
}
