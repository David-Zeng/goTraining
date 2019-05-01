//////////////////////////////////////////////////////////////////////////////////////////////
// name: 09_slice
//
//
//////////////////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0:len(s)])

	var s1 []int
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)

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



}
