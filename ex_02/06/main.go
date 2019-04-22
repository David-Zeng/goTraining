package main

import "fmt"

const (
	_       = iota
	nextYr1 = iota + 2019
	nextYr2 = iota + 2019
	nextYr3 = iota + 2019
	nextYr4 = iota + 2019
)

func main() {

	fmt.Println(nextYr1)
	fmt.Println(nextYr2)
	fmt.Println(nextYr3)
	fmt.Println(nextYr4)

}
