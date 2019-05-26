package main

import (
	"encoding/json"
	"fmt"
	"github.com/David-Zeng/goTraining/myFunctions"
)

type person struct {
	First string
	Last  string
	Age   int
}

type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	fmt.Println("json")

	fmt.Println("@@@@@@@@@@@@1")
	p1 := person{
		First: "jones",
		Last:  "bond",
		Age:   38,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	myFunctions.ObjectDescribe(p1)
	myFunctions.ObjectDescribe(p2)

	persons := []person{p1, p2}
	myFunctions.ObjectDescribe(persons)

	fmt.Println("@@@@@@@@@@@@2")
	bs, err := json.Marshal(persons)

	myFunctions.ObjectDescribe(bs)
	myFunctions.ObjectDescribe(err)

	fmt.Println(string(bs))

	fmt.Println("@@@@@@@@@@@@3")
	var p3 []person2

	json.Unmarshal(bs, &p3)

	myFunctions.ObjectDescribe(p3)

}
