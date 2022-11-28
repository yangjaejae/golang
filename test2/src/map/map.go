package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)

	m["aaa"] = "ccc"
	fmt.Println(m["aaa"])

	mm := make(map[string]string)
	mm["aaa"] = "ccc"
	fmt.Println(mm["aaa"])

	m2 := make(map[int]int)

	m2[3] = 0
	m2[5] = 10

	fmt.Println(m2[3])
	fmt.Println(m2[4])

	v, ok := m2[4]

	// map 리턴 값의 두번째 값은 값의 유무를 나타내는 boolean 값이 나옴
	fmt.Println(v, ok)

	delete(m2, 5)

	v, ok = m2[5]
	fmt.Println(v, ok)

	for key, value := range m2 {
		fmt.Println(key, "=", value)
	}
}
