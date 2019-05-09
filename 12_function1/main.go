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

	fmt.Println(IntSumFunc(1,2,3,4,5,6,7,8,9))

	fmt.Println(getFactorial(5))

}

// simple function
func doSomething() {
	fmt.Println("do something")
}

// sum of ints
func IntSumFunc(nums ...int) int {
	fmt.Printf("%T\n", nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
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

