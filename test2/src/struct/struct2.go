package main

import "fmt"

type Student struct {
	name  string
	class int

	score Score
}

type Score struct {
	name  string
	grade string
}

func (s Student) ViewScore() {
	fmt.Println(s.score)
}

func (s Student) InputScore(name string, grade string) {
	s.score.name = name
	s.score.grade = grade
}

func main() {
	var s Student
	s.name = "tom"
	s.class = 1

	s.score.name = "math"
	s.score.grade = "c"

	s.ViewScore()

	s.InputScore("science", "10")
	s.ViewScore()
}
