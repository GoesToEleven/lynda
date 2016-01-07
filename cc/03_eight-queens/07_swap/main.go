package main

import "fmt"

func main() {
	options := []string{"A", "B"}
	fmt.Println(options)
	swap(options, 0, 1)
	fmt.Println(options)
}

func swap(a []string, x, y int) {
	a[x], a[y] = a[y], a[x]
}
