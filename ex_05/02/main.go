package main

import "fmt"

func main() {
	fmt.Println("map of person struct")

	type person struct {
		first string
		last string
		favIceCream []string
	}

	p1 := person{
		first: "James",
		last:  "bond",
		favIceCream: []string{
			"Rum Raisin",
			"Cherry",
			"pistachio"},
	}

	p2 := person{
		first: "miss",
		last:  "moneypenny",
		favIceCream: []string{
			"chocolate",
			"double'o'fudge",
			"caramel"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
