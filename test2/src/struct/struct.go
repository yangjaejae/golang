package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintName() {
	fmt.Println(p.name)
}

func main() {
	var p Person
	p1 := Person{"john", 12}
	p2 := Person{name: "james", age: 40}
	p3 := Person{name: "tom"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	p.name = "smith"
	p.age = 24
	fmt.Println(p)

	p.PrintName()
}
