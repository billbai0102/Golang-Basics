package main

import "fmt"

type person struct {
	name string
	age int
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main(){
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 21})

	fmt.Println(person{name: "Frederick"})

	fmt.Println(&person{name: "Anne", age: 40})

	fmt.Println(NewPerson("John"))

	s := person{name:"Sean", age: 50}
	fmt.Println(s.name)

	sptr := &s
	fmt.Println(sptr.age)

	sptr.age = 51
	fmt.Println(sptr.age)
}