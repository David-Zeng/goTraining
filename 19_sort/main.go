package main

import (
	"fmt"
	"github.com/David-Zeng/goTraining/myFunctions"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	fmt.Println("sort pkg")

	fmt.Println("@@@@@@@@@@@@@@@@@1")
	s := []int{3, 5, 7, 9, 2, 1, 0}

	myFunctions.ObjectDescribe(s)

	sort.Ints(s)
	myFunctions.ObjectDescribe(s)

	fmt.Println("@@@@@@@@@@@@@@@@@2")
	p1 := person{
		Name: "James Bond",
		Age:  38,
	}
	myFunctions.ObjectDescribe(p1)

	p2 := person{
		Name: "Money Penny",
		Age:  42,
	}
	myFunctions.ObjectDescribe(p2)

	fmt.Println("@@@@@@@@@@@@@@@@@3")
	people := []person{p1, p2}

	myFunctions.ObjectDescribe(people)
	myFunctions.ObjectDescribe(ByAge(people))

	fmt.Println(ByAge(people).Len())

	ByAge(people).Swap(0, 1)
	myFunctions.ObjectDescribe(ByAge(people))

	sort.Sort(ByAge(people))
	myFunctions.ObjectDescribe(ByAge(people))
}
