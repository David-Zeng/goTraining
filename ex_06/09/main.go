package main

import "fmt"

func main() {
	fmt.Println("function callback")

	fmt.Println(sum([]int{2, 4, 6, 8}...))

	fmt.Println(even(sum, []int{1, 2, 3}...))
}

func sum(nums ...int) int {
	tot := 0
	for _, v := range nums {
		tot += v
	}

	return tot
}

func even(f func(nums ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}



