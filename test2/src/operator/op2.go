package main

import "fmt"

func main() {
	a := 3
	if a == 3 || a == 4 {
		fmt.Println("참")
	} else if a != 3 && a == 4 {
		fmt.Println("거짓")
	} else {
		fmt.Println("몰라")
	}

}
