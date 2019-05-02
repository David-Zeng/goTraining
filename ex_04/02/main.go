////////////////////////////////////////////////////////////////////////////////
// name: ex_04/02
// requirement:  1. create a slice hold 10 values of type int
//
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

func main() {
	fmt.Println("@@@part1")
	s := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println(s)

	for _, v := range s {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", s)
}
