package main

import (
	"fmt"
	"github.com/David-Zeng/goTraining/myFunctions"
)

// person type
type person struct {
	first string
	last  string
	phone string
}

func (p person) speak() {
	fmt.Println("This is", p.first, p.last, "talking", "- a person")
}

// secretAgent
type secretAgent struct {
	person person
	ltk    bool
}

func (sa secretAgent) speak() {
	fmt.Println("This is", sa.person.first, sa.person.last, "speaking", "- a speicalAgent")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("this is a person type: ", h)
	case secretAgent:
		fmt.Println("this is a secretAgent type: ", h)
	default:
		fmt.Println("other type:", h)
	}
}

func main() {

	fmt.Println("struct function & interface")

	fmt.Println("@@@@@@@@@@1: secretAgent type")
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa1.speak()

	sa2 := secretAgent{
		person: person{
			first: "Money",
			last:  "Penny",
		},
		ltk: true,
	}

	fmt.Println(sa1.person.first)
	fmt.Println(sa1.person.last)
	fmt.Println(sa1.ltk)

	sa2.speak()
	sa2.person.speak()

	fmt.Println("@@@@@@@@@@2: person type")
	p1 := person{
		first: "David",
		last:  "Zeng",
	}

	fmt.Println(p1)
	p1.speak()

	fmt.Println("@@@@@@@@@@3: human interface")
	bar(p1)
	bar(sa1)
	bar(sa2)

	fmt.Println("@@@@@@@@@@4: nil interface")
	var i interface{}
	i = "string"
	myFunctions.ObjectDescribe(i)
	i = 41.5
	myFunctions.ObjectDescribe(i)
	i = []int{1, 2, 3, 4, 5, 6, 7}
	myFunctions.ObjectDescribe(i)
}
