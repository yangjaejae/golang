package main

import "fmt"

func main() {
	i := 0

	for i < 10 {
		fmt.Println(i)

		i += 1
	}
	fmt.Println("##")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("##")
	j := 0

	for {
		if j == 5 {
			j++
			continue
		}
		if j == 6 {
			break
		}
		fmt.Println(j)
		j++
	}
}
