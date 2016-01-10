package main

import "fmt"

func main() {
	buckets := make([]int, 1)
	fmt.Println(buckets[0])
	buckets[0] = 42
	fmt.Println(buckets[0])
	buckets[0]++
	fmt.Println(buckets[0])
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
