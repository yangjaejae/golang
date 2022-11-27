package main

func main() {
	i := 0

L1:
	for {
		println("first in")
		if i == 0 {
			println("break")
			break L1
		}
	}

	println("OK")
}
