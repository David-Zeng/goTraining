package main

import "fmt"

func main() {

	slice1 := []string{"James", "Bond", "Shaken, not stirred"}
	slice2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	slice := [][]string{slice1, slice2}

	fmt.Println(slice)

	for i, v := range slice {
		for i2, v2 := range v {
			fmt.Println(i, i2, v2)
		}
	}

}
