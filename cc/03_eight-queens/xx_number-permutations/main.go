package main

import "fmt"

func main() {
	str := "ABCD"
	n := len(str)
	fmt.Println("Number of permutations", factorial(n))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// 1 * 1 = 1
// 2 * 1 * 1 = 2
// 3 * 2 * 1 * 1 = 6
// 4 * 3 * 2 * 1 * 1 = 24
// 5 * 4 * 3 * 2 * 1 * 1 = 120
// 6 * 5 * 4 * 3 * 2 * 1 * 1 = 720
// 7 * 6 * 5 * 4 * 3 * 2 * 1 * 1 = 5040