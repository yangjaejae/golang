package main

import "fmt"

func main() {
	// 초기에 capacity가 가득찬 상태로 생성되었기 때문에
	// append를 하여 다른 변수에 값을 할당하게 될 경우
	// capacity를 늘린 새로운 메모리 주소에 slice를 생성하여
	// 완전히 다른 slice가 탄생하게 된다
	a := []int{1, 2}
	b := append(a, 3)

	fmt.Printf("%p, %v \n", a, a)
	fmt.Printf("%p, %v \n", b, b)

	// 초기에 catapity 남는 상태로 생성되었기 때문에
	// append를 하여 변수에 값을 할당하게 될 경우
	// capacity를 늘리지 않아도 되기 때문에 같은 slice 주소를
	// 가지게 된다.
	c := make([]int, 2, 10)
	d := append(c, 3)
	fmt.Printf("%p, %v \n", c, c)
	fmt.Printf("%p, %v \n", d, d)
}

// 결과 값
// 0xc0000140c0, [1 2]
// 0xc0000220a0, [1 2 3]
// 0xc00001c0f0, [0 0]
// 0xc00001c0f0, [0 0 3]
