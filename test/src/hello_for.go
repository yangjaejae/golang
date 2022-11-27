package main

func main() {

	for i := 0; i < 5; i++ {
		println("hello")
	}

	n := 10
	for n < 20 {
		println(n)
		n++
	}

	// for {
	// 	println("무한 루프")
	// }

	names := []string{"yang", "jae", "ho"}
	for i, v := range names {
		println(i, " : ", v)
	}
}
