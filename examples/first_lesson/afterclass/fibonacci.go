package main

func main() {
	println(fibonacciIter(0))
	println(fibonacciIter(1))
	println(fibonacciIter(2))
	println(fibonacciIter(3))
	println(fibonacciIter(4))
	println(fibonacciIter(6))

	println("----")
	println(fibonacci(0))
	println(fibonacci(1))
	println(fibonacci(2))
	println(fibonacci(3))
	println(fibonacci(4))
	println(fibonacci(6))
}

func fibonacciIter(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := 0
	b := 1
	for i := 2; i < n; i++ {
		t := b
		b = b + a
		a = t
	}
	return b + a
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
