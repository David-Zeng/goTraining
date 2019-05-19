package main

import (
	"fmt"
	"github.com/David-Zeng/goTraining/myFunctions"
)

func main() {

	fmt.Println("pointer ex 01")

	a := 1

	myFunctions.ObjectDescribe(&a)
}
