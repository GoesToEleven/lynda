package main

import "fmt"

func main() {
	f := fibonacci(4)
	fmt.Println(f)
}

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
