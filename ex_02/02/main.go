package main

import (
	"fmt"
)

func main() {

	x := ("a" != "b")
	y := ("b" > "a")
	z := ("a" > "c")

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", z, z)

	fmt.Printf("%v\n", []byte("david"))
}
