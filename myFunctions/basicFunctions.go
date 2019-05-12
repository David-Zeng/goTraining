package myFunctions

import "fmt"

func ObjectDescribe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
