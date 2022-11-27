package main

func main() {
	var a = 1
	for a < 15 {
		if a == 5 {
			println("first if: ", a)
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			println("break if: ", a)
			break //루프 빠져나옴
		}

		if a < 5 {
			println("under 5: ", a)
			goto END
		}
	}
	if a == 11 {
		println("goto if: ", a)
		goto END //goto 사용예
	}
	println("그냥: ", a)

END:
	println("End")
}
