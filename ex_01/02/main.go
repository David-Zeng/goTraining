////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// name: 
// purposer:
// history:
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

var (
	x int    //default value 0
	y string //default value ''
	z bool   //default value false 
)

func main() {
	
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", z, z)
}
