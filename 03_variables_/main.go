package main

import (
	"fmt"
)

var (
	z = 21
)

func main() {

	//var (
	var x int
	x = 42
	//)

	fmt.Println("value:", x)
	fmt.Printf("%T\n", x)

	fmt.Println("value:", z)
	fmt.Printf("%T\n", z)

}


