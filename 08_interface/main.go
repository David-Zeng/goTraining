package main

import "fmt"

func main() {
	var i interface{} = 23
	fmt.Printf("%d\t%v\t%#v\t%T\n", i, i, i, i)
}
