package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++

		if i == 5 {
			break
		}
	}

	i = 1
	for i <= 10 {
		if i == 5 {
			continue
		}
		fmt.Println(i)
		i++
	}

	switch_func()
}

func switch_func() {
	switch {
	case false:
		fmt.Println("A")
	case 1 == 1:
		fmt.Println("B")
		//fallthrough
	case 2 == 2:
		fmt.Println("C")
	default:
		fmt.Println("other")
	}

}
