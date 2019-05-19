package main

import "fmt"

func main() {
	fmt.Println("requirement: assign a function to a variable")

	helloWorldStr := func() {
		fmt.Println("hello world")
	}

	helloWorldStr()
}
