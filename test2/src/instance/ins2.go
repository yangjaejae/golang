package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func (t *Student) SetName(newName string) {
	t.name = newName
}

func (t *Student) SetAge(newAge int) {
	t.age = newAge
}

func (t *Student) PrintStudent() {
	fmt.Println(t)
}

func main() {
	var a *Student
	a = &Student{"aaa", 20, 10}
	// a := Student{"aaa", 20, 10}

	a.SetName("bbb")
	fmt.Println(a)

	a.SetAge(30)
	fmt.Println(a)

	a.PrintStudent()
}
