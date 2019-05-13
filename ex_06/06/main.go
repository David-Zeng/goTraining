package main

import "fmt"

func main() {
	fmt.Println("requirement: build an anonymous function")

	func (s string) {
		fmt.Println(s)
	}("Hello world, anonymous function!")

}


