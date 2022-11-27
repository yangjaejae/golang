package main

func main() {
	var a [3]int //정수형 3개 요소를 갖는 배열 a 선언
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1]) // 2 출력

	string_arr := []string{"a", "b", "c"}
	println(string_arr)
	for i, v := range string_arr {
		println(i, " ", v)
	}

	for i := 0; i < len(string_arr); i++ {
		println(string_arr[i])
	}
}
