////////////////////////////////////////////////////////////////////////////////
// name: 12_function1
//
//
//
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

func main() {

	doSomething()

	fmt.Println(getFactorial(5))
}

// simple function
func doSomething() {
	fmt.Println("do something")
}

// recursive
// n! = n×(n-1)! where n >0
func getFactorial(num int) int {
	if num > 1 {
		return num * getFactorial(num-1)
	} else {
		return 1 // 1! == 1
	}
}
