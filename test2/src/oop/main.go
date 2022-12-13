package main

import "fmt"

type Bread struct {
	val string
}

type Jam struct {
}

func (b *Bread) PutJam(jam *Jam) {
	b.val += jam.GetVal()
}

func (j *Jam) GetVal() string {
	return "jam"
}

func (b *Bread) String() string {
	return b.val
}

func main() {
	bread := &Bread{val: "bread"}
	jam := &Jam{}

	bread.PutJam(jam)

	fmt.Println(bread)
	// 객체를 출력했는데 객체에 "String"이라는 메서드가 있는 경우 그 메서드를 실행
}
