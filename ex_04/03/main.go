////////////////////////////////////////////////////////////////////////////////
// name: ex_04/03
// requirement:  1. create the following slices
//					[42 43 44 45 46]
//					[47 48 49 50 51]
//					[44 45 46 47 48]
//					[43 44 45 46 47]
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"github.com/David-Zeng/goTraining/myFunctions"
)

func main() {

	var s []int
	for i := 42; i <= 51; i += 1 {
		fmt.Println(i)
		s = append(s, i)
		fmt.Println(i, "s len: ", len(s))
		fmt.Println(i, "s cap: ", cap(s))
	}

	fmt.Println(s)

	fmt.Println(myFunctions.IntSliceSliceWithLengthFunc(s, myFunctions.IntFindSliceElementPositionFunc(s, 42), 5))
	fmt.Println(myFunctions.IntSliceSliceWithLengthFunc(s, myFunctions.IntFindSliceElementPositionFunc(s, 47), 5))
	fmt.Println(myFunctions.IntSliceSliceWithLengthFunc(s, myFunctions.IntFindSliceElementPositionFunc(s, 44), 5))
	fmt.Println(myFunctions.IntSliceSliceWithLengthFunc(s, myFunctions.IntFindSliceElementPositionFunc(s, 43), 5))

}


