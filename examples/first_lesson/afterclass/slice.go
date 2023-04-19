package main

func main() {
	s := []int{1, 2, 4, 7}
	// 结果应该是 5, 1, 2, 4, 7
	s = Add(s, 0, 5)
	// 打印s的内容
	PrintSlice(s)

	// 结果应该是5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)
	PrintSlice(s)

	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)
	PrintSlice(s)

	// 结果应该是5, 9, 2, 4, 7, 13
	s = Delete(s, 2)
	PrintSlice(s)

	// 结果应该是9, 2, 4, 7, 13
	s = Delete(s, 0)
	PrintSlice(s)

	// 结果应该是9, 2, 4, 7
	s = Delete(s, 4)
	PrintSlice(s)

}

func PrintSlice(s []int) {
	for i := 0; i < len(s); i++ {
		print(s[i], ", ")
	}
	println()
}

func Add(s []int, index int, value int) []int {
	s = append(s, 0)
	for i := len(s) - 1; i > index; i-- {
		s[i] = s[i-1]
	}
	s[index] = value
	return s
}

func Delete(s []int, index int) []int {
	for i := index; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	s = s[:len(s)-1]
	return s
}
