package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello World!")

	foo()

	for x := 0; x < 10; x ++ {
		if x%2 == 0 {
			fmt.Println(x)
		}
	}

	bar()

	var str1 string = "hahahaha"
	fmt.Println(reflect.TypeOf(str1).Kind())

	n, _ := fmt.Println("Hohoho", true)
	fmt.Println(n)
}

func foo() {
	fmt.Println("I am in foo")
}


func bar() {
	fmt.Println("and then we exited")
}



