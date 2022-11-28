package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintScore() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) inputScore(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var s Student = Student{name: "john", age: 22, class: "math", grade: "A"}

	s.inputScore("science", "c")
	s.PrintScore()
}
