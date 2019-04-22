package main

import "fmt"

const (
	a = 42
	b = 42.78
	c = "James Bond"

	d int     = 42
	e float64 = 42.78
	f string  = "James Bond"
)

func main() {
	fmt.Printf("a: %T %v\n", a, a)
	fmt.Printf("b: %T %v\n", b, b)
	fmt.Printf("c: %T %v\n", c, b)

	fmt.Printf("d: %T %v\n", d, d)
	fmt.Printf("e: %T %v\n", e, e)
	fmt.Printf("f: %T %v\n", f, f)
}
