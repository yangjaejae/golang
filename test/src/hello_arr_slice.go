package main

import "fmt"

func main() {
	var a []int        //슬라이스 변수 선언
	a = []int{1, 2, 3} //슬라이스에 리터럴값 지정
	a[1] = 10
	fmt.Println(a) // [1, 10, 3]출력

	s := make([]int, 5, 10)
	println(len(s), cap(s)) // len 5, cap 10

	ss := []int{0, 1, 2, 3, 4, 5}
	ss = ss[2:5]
	fmt.Println(ss) //2,3,4 출력
	ss = ss[2:]
	fmt.Println(ss) //4 출력

	sss := []int{0, 1}

	// 하나 확장
	sss = append(sss, 2) // 0, 1, 2
	// 복수 요소들 확장
	sss = append(sss, 3, 4, 5) // 0,1,2,3,4,5

	fmt.Println(sss)

	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...)
	//sliceA = append(sliceA, 4, 5, 6)

	fmt.Println(sliceA) // [1 2 3 4 5 6] 출력

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)               // [0 1 2 ] 출력
	println(len(target), cap(target)) // 3, 6 출력
}
