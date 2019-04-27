package main

import (
	"fmt"
	"time"
)

func main() {
	//for condition + incremental
	for i := 65; i <= 90; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Printf("%#U\n", i)
		}
	}

	//for condition
	currentTime := time.Now()
	a := 1979
	for a <= currentTime.Year() {
		fmt.Println(a)
		a++
	}

	//for
	a = 1979
	for {
		if a > currentTime.Year() {
			break
		}
		fmt.Println(a)
		a++
	}
}
