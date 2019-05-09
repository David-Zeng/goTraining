package main

import "fmt"


type person struct {
	first string
	last string
	phone string
}

type specialAgen struct {
	person person
	ltk bool
}

func (sa specialAgen) speak() {
	fmt.Println("This is", sa.person.first, sa.person.last, "speaking")
}

func (p person) talk() {
	fmt.Println("This is", p.first, p.last, "talking")
}

func main() {

	fmt.Println("struct function & interface")

	sa1 := specialAgen{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}
	sa1.speak()

	sa2 := specialAgen{
		person: person{
			first: "Money",
			last: "Penny",
		},
		ltk: true,
	}
	sa2.speak()

	sa2.person.talk()

	fmt.Println(sa1.person.first)
	fmt.Println(sa1.person.last)
	fmt.Println(sa1.ltk)

}



