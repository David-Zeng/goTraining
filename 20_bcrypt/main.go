package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("bcrypt pkg")

	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPassword1 := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword1))
	fmt.Println(err)
	if err != nil {
		fmt.Println("you can't login")
		return
	}
	fmt.Println("you are logged in")
}


