package main

import (
	"fmt"
)

type (
	hotdog int
)

var (
	y = 42
	x int
	a float64

	z = "David Z"
	str1 = `line1
		    line2
		    line3
			`
	b hotdog

	c [4]int
)

func main() {

	fmt.Println("y value:", y)
	fmt.Printf("y type %T\n:", y)

	fmt.Println("z value:", z)
	fmt.Printf("z type %T\n:", z)

	fmt.Println("str1 value:", str1)
	fmt.Printf("str1 type %T\n:", str1)

	y = int(b)
	fmt.Println("y value:", y)
	fmt.Printf("y type %T\n:", y)

	x = 2
	fmt.Println("x value:", x)
	fmt.Printf("x type %T\n:", x)

	a = 109180983097123097
	fmt.Println("a value:", a)
	fmt.Printf("a type %T\n:", a)

	bs := []byte(z)
	fmt.Println("bs value:", bs)
	fmt.Printf("bs type %T\n:", bs)

	var s []int
	//s := make([]int, 10, 10)
	fmt.Printf("s len=%d cap=%d %v\n", len(s), cap(s), s)
	for i := 0; i <= 10; i++ {
		s = append(s, i)
		fmt.Printf("s len=%d cap=%d %v\n", len(s), cap(s), s)
	}

	s2 := []int{0, 1, 2}
	fmt.Printf("s2 len=%d cap=%d %v\n", len(s2), cap(s2), s2)

	fmt.Println("value:", c)
	fmt.Printf("c type %T\n:", c)
	fmt.Printf("c type %T\n:", c[:])

}




