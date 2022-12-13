package main

import "fmt"

type Bread struct {
	val string
}

type Jam struct {
	opened bool
}

type SpoonOfJam struct {
}

type Sandwitch struct {
	val string
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num)
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}

func OpenJam(jam *Jam) {
	jam.opened = true
}

func GetOneSpoon(_ *Jam) *SpoonOfJam {
	return &SpoonOfJam{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfJam) {
	bread.val += " + Jam"
}

func MakeSandWitch(breads []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + " + "
	}
	return sandwitch
}

func main() {
	// 1. 빵 두개를 꺼낸다
	breads := GetBreads(2)

	jam := &Jam{}
	// 2. 딸기잼 뚜껑을 연다
	OpenJam(jam)
	// 3. 딸기잼 한스푼 뜬다
	spoon := GetOneSpoon(jam)
	// 4. 딸기잼을 빵에 바른다.
	PutJamOnBread(breads[0], spoon)
	// 5. 빵을 덮는다.
	sandwitch := MakeSandWitch(breads)
	// 6. 완성
	fmt.Println(sandwitch.val)
}
