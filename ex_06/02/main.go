package main

import "fmt"

func main() {

	fmt.Println("@@@@@@@@@@@@1")
	fmt.Println(bar())

	fmt.Println("@@@@@@@@@@@@2")
	fmt.Println("int... sum: ")
	s := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(intSumFunc(s...))

	fmt.Println("slice sum: ")
	fmt.Println(sliceSumFunc(s))

	ss := []interface{}{"A", "B", "C"}
	interfaceSumFunc(ss...)

	si := []interface{}{1, 2, 3}
	interfaceSumFunc(si...)

}

func bar() (int, string) {
	return 1984, "Big Brother is Watching"
}

func intSumFunc(ints ...int) int {
	tot := 0
	for _, v := range ints {
		tot += v
	}
	return tot
}

func sliceSumFunc(si []int) int {
	tot := 0
	for _, v := range si {
		tot += v
	}
	return tot
}

func interfaceSumFunc(is ...interface{}) int {
	//tot := 0
	for _, v := range is {
		s, ok := v.(int)
		fmt.Println(v)
		fmt.Println(s)
		fmt.Println(ok)
	}
	return 0
}


