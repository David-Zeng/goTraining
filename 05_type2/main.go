package main

import "fmt"

func main() {

	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}

	fmt.Printf("%T %v\n", d, d)

	var f []byte

	copy(d, f)
	fmt.Printf("%T %v\n", d, d)

	//f[2] = 'z'
	fmt.Printf("%T %v\n", d, d)
	fmt.Printf("%T %v\n", f, f)
}
