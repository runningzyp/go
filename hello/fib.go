package main

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		return fib(n-2) + fib(n-1)
	}

}

func main() {
	for i := 1; i < 10; i++ {
		print(" ", fib(i))
	}
}
