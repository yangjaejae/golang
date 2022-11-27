package main

func main() {
	total := sum(1, 7, 3, 5, 9)
	println(total)

	count, tot := _sum(1, 2, 3, 4, 5)
	println("count: ", count)
	println("tot: ", tot)

	cnt, tota := __sum(1, 2, 3, 4, 5)
	println("cnt: ", cnt)
	println("tota: ", tota)
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func _sum(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func __sum(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
