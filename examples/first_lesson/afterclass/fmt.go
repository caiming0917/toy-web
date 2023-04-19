package main

import "fmt"

func main() {
	println(printNumWith2(1.234))
	println(printNumWith2(1.2))
	println(printNumWith2(1.0))
	println(printNumWith2(1.000))

	println(printBytes([]byte("Hello World!")))
}

// 输出两位小数
func printNumWith2(float642 float64) string {
	return fmt.Sprintf("%.2f", float642)
}

func printBytes(data []byte) string {
	return fmt.Sprintf("%s", data)
}
