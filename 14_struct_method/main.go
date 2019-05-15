package main

import "fmt"

/*
In Go, values can be of more than one type. An interface allows a value to be of more than one type.
We create an interface using this syntax: “keyword identifier type” so for an interface it would be:
“type human interface” We then define which method(s) a type must have to implement that interface.
If a TYPE has the required methods, which could be none (the empty interface denoted by interface{}),
then that TYPE implicitly implements the interface and is also of that interface type.
In Go, values can be of more than one type.


*/

type person struct {
	first string
	last string
	phone string
}

type secretAgent struct {
	person person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("This is", sa.person.first, sa.person.last, "speaking")
}

func (p person) speak() {
	fmt.Println("This is", p.first, p.last, "talking")
}

type human interface {
	speak()
}

// function bar can take person type or secretAgent type (because both of types are human type as well)
func bar(h human) {

	//h can be person type or secretAgent
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).person.first)
	}
	fmt.Println("I was passed into bar", h)
}


func main() {

	fmt.Println("struct function & interface")

	fmt.Println("@@@@@@@@@@1")
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}
	sa1.speak()

	sa2 := secretAgent{
		person: person{
			first: "Money",
			last: "Penny",
		},
		ltk: true,
	}
	sa2.speak()

	sa2.person.speak()


	fmt.Println("@@@@@@@@@@2")
	p1 := person{
		first: "David",
		last: "Zeng",
	}

	fmt.Println(sa1.person.first)
	fmt.Println(sa1.person.last)
	fmt.Println(sa1.ltk)

	fmt.Println(p1)
	p1.speak()

	fmt.Printf("%T\n", sa1)
	fmt.Printf("%T\n", sa2)


	fmt.Println("@@@@@@@@@@3")
	bar(sa1)
}



