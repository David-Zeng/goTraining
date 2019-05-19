package main

import "fmt"

func main() {
	fmt.Println("function output a function")

	fmt.Println(f()())

}

func f() func() int {
	return func() int {
		return 100
	}
}
