package main

import "fmt"

func main() {
	var a []int               // 초기값이 없는 동적 배열 생성
	b := []int{1, 2, 3, 4, 5} // 초기값을 지정한 동적 배열 생성
	c := make([]int, 0, 8)    // 초기값이 없이 초기 length와 capacity를 지정하여 동적 배열 생성

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("len(a) = %d\n", cap(a))
	fmt.Printf("len(a) = %d\n", len(b))
	fmt.Printf("len(a) = %d\n", cap(b))
	fmt.Printf("len(a) = %d\n", len(c))
	fmt.Printf("len(a) = %d\n", cap(c))
}
