////////////////////////////////////////////////////////////////////////////////
// name: 11_map
//
//
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

func main() {

	fmt.Println("@1: part1")
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["David"])

	if v, ok := m["David"]; ok {
		fmt.Println("David exists in map m:", v)
	} else {
		fmt.Println("David doesn't exists in map m:")
	}

	fmt.Println("@1: part2")
	n := map[string]int{
		"Joanne": 41,
		"Olivia": 8,
	}

	if g, err := n["David"]; !err {
		n["David"] = 40
		fmt.Println("Added David into map n")
	} else {
		fmt.Println(g)
	}

	fmt.Println(n["David"])

}
