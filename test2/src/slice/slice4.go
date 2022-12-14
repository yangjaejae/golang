package main

import "fmt"

func main() {
	var s []int
	var t []int

	s = make([]int, 3)

	s[0] = 100
	s[1] = 200
	s[2] = 300

	fmt.Println(s, len(s), cap(s))

	s = append(s, 400)
	fmt.Println(s, len(s), cap(s))

	t = append(s, 400)
	fmt.Println(t, len(t), cap(t))

	var u []int
	u = append(t, 500)

	u[0] = 9999

	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))
	fmt.Println(u, len(u), cap(u))
}
