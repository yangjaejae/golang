package main

import "fmt"

type SpoonOfJam interface {
	String() string
}

type Jam interface {
	// GetOneSpoon() *SpoonOfStrawberryJam
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	val string
}

type StrawberryJam struct {
}

type SpoonOfStrawberryJam struct {
}

type OrangeJam struct {
}

type SpoonOfOrangeJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return " + strawberry"
}

func (s *SpoonOfOrangeJam) String() string {
	return " + Orange"
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}

func main() {
	bread := &Bread{}
	// jam := &StrawberryJam{}
	jam := &OrangeJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}
