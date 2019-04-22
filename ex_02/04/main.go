package main

import (
	"fmt"
)

func main() {

	a := 2
	fmt.Printf("%v %T\n", a, a)

	b := a << 1
	fmt.Printf("%v %T\n", b, b)

}
