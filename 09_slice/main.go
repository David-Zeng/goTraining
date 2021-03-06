//////////////////////////////////////////////////////////////////////////////////////////////
// name: 09_slice
//
//
//////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"github.com/David-Zeng/goTraining/myFunctions"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0:len(s)])
	fmt.Printf("%T", s)

	var s1 []int
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)
	fmt.Printf("%T", s1)

	for i, v := range s {
		fmt.Println(i)
		fmt.Println(v)

		s1 = append(s1, v)
	}
	s1 = append(s1, s...)
	fmt.Printf("s1 slice: %v\n", s1)

	for j := 0; j < len(s); j++ {
		fmt.Println(j)
		fmt.Printf("s slice value: %v\n", s[j])
	}
	cs := append(s[:3], s[4:]...)
	fmt.Println(cs)

	myFunctions.IntSliceDelFunc(s1, 3)
	myFunctions.IntSliceDelFunc(s, 3)

}
