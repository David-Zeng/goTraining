package main

import (
	"fmt"
	"github.com/David-Zeng/goTraining/myFunctions"
)

func main() {
	fmt.Println("ex_04 05")
	fmt.Println("requirement: start with slice = {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}")
	fmt.Println("requirement: return slice  {42, 43, 44, 45, 51}")

	s := []int{42, 43, 44, 45, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println("@@@@@@@@@1")
	fmt.Println(s)
	fmt.Println(myFunctions.IntSliceElementExistFunc(s, 100))
	fmt.Println(s)
	fmt.Println(myFunctions.IntSliceElementExistFunc(s, 45))
	fmt.Println(s)

	fmt.Println("@@@@@@@@@2")
	fmt.Println("remove slice by position testing:")
	fmt.Println(s)
	fmt.Println(myFunctions.IntSliceDelByPositionMultiFunc(s, 4, 5, 6))
	fmt.Println(s)

	fmt.Println("@@@@@@@@@3")
	s = []int{42, 43, 44, 45, 45, 46, 47, 48, 49, 50, 51, 51, 51}
	fmt.Println("deduplicated test")
	fmt.Println(s)
	fmt.Println(myFunctions.IntSliceDeduplicateFunc(s))

	fmt.Println("@@@@@@@@@4")
	fmt.Println(myFunctions.IntSlicesDiffFunc(s, []int{43, 50}))

}
