package main

import (
"encoding/json"
"fmt"
"github.com/David-Zeng/goTraining/myFunctions"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {

	fmt.Println("@@@@@@@@@@@@@1")
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := Person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	myFunctions.ObjectDescribe(p1)
	myFunctions.ObjectDescribe(p2)

	persons := []Person{p1, p2}
	myFunctions.ObjectDescribe(persons)

	fmt.Println("@@@@@@@@@@@@2")
	bs, err := json.Marshal(persons)

	myFunctions.ObjectDescribe(bs)
	myFunctions.ObjectDescribe(err)

	fmt.Println(string(bs))

	people := []Person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	myFunctions.ObjectDescribe(bs)

	myFunctions.ObjectDescribe(err)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	myFunctions.ObjectDescribe(string(bs))

	fmt.Println("@@@@@@@@@@@@@2")
	fmt.Println(string(91))
	fmt.Println(string(123))
	fmt.Println(string(123))


	fmt.Println("@@@@@@@@@@@@@3")
	jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	var v interface{}
	json.Unmarshal(jsonData, &v)
	myFunctions.ObjectDescribe(v)

	data := v.(map[string]interface{})
	myFunctions.ObjectDescribe(data)

	a, _ := json.Marshal(data)

	myFunctions.ObjectDescribe(string(a))

}

