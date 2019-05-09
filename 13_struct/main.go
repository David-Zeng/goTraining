package main

import "fmt"

func main() {
	fmt.Println("struct")

	fmt.Println("@@@@@1")
	type person struct {
		first string
		last string
	}

	me := person{
		first: "David",
		last: "Zeng",
	}

	fmt.Println(me)
	fmt.Println(me.first)
	fmt.Println(me.last)


	fmt.Println("@@@@@2")
	type speicalAgent struct {
		person
		ltk bool
	}

	sa := speicalAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}

	fmt.Println(sa)
	fmt.Println(sa.first)
	fmt.Println(sa.last)
	fmt.Println(sa.ltk)

	fmt.Println("@@@@@3")

	sa2 := struct {
		first string
		last string
		ltk bool
	} {
		first: "James",
		last: "Bond",
		ltk: true,
	}

	fmt.Println(sa2)

}

