package main

import "fmt"

func main() {
	x := 33

	switch { // switch문에 변수가 아무것도 들어가지 않을 경우에는 "true"가 들어감
	case x > 40:
		fmt.Println("x는 40보다 크다")
	case x < 20:
		fmt.Println("x는 20보다 작다")
	case x > 30:
		fmt.Println("x는 30보다 크다") // 먼저 걸리는 쪽으로 하고 끝남 break 사용 안해도 됨
	case x == 33:
		fmt.Println("x는 33")
	}
}
