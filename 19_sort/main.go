package main

import (
	"fmt"
	"github.com/David-Zeng/goTraining/myFunctions"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByName []Person
func (bn ByName) Len() int {
	return len(bn)
}
func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}
func (bn ByName) Less(i, j int) bool {
	return bn[i].Name < bn[j].Name
}


func main() {
	fmt.Println("sort pkg")

	fmt.Println("@@@@@@@@@@@@@@@@@1")
	s := []int{3, 5, 7, 9, 2, 1, 0}

	myFunctions.ObjectDescribe(s)

	sort.Ints(s)
	myFunctions.ObjectDescribe(s)

	fmt.Println("@@@@@@@@@@@@@@@@@2")
	p1 := Person{
		Name: "James Bond",
		Age:  38,
	}
	myFunctions.ObjectDescribe(p1)

	p2 := Person{
		Name: "Money Penny",
		Age:  42,
	}
	myFunctions.ObjectDescribe(p2)

	fmt.Println("@@@@@@@@@@@@@@@@@3")
	people := []Person{p1, p2}

	myFunctions.ObjectDescribe(people)
	myFunctions.ObjectDescribe(ByAge(people))

	fmt.Println(ByAge(people).Len())

	ByAge(people).Swap(0, 1)
	myFunctions.ObjectDescribe(ByAge(people))

	sort.Sort(ByAge(people))
	myFunctions.ObjectDescribe(ByAge(people))


	fmt.Println("@@@@@@@@@@@@@@@@@4")
	myFunctions.ObjectDescribe(ByName(people))

}
