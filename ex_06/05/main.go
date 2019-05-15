<<<<<<< Updated upstream
package main

import (
	"fmt"
	"math"
)

type square struct {
	l float64
	w float64
}
func (s square) area() float64 {
	return s.l * s.w
}


type circle struct {
	r float64
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

type shape interface {
	area() float64
}

func info(si shape)  {
	fmt.Println(si.area())
}


func main() {
	fmt.Println("square & circle area")

	s1 := square{
		l: 5.0,
		w: 5.0,
	}

	c1 := circle{
		r: 5.0,
	}

	info(s1)
	info(c1)
}
=======
package _5
>>>>>>> Stashed changes
