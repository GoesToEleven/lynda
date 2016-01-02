package main

import "fmt"

func main() {
	options := []int{1, 2, 3}
	perm := permutations(options)
	for _, v := range perm {
		fmt.Println(v)
	}
	fmt.Println("Number of answers:", len(perm))
}

func permutations(options []int) [][]int {
	return permute(options, 0)
}

func swap(a []int, x, y int) {
	a[x], a[y] = a[y], a[x]
}

func permute(options []int, start int) [][]int {
	answers := [][]int{}
	end := len(options) - 1
	if start == end {
		ans := make([]int, len(options))
		copy(ans, options)
		answers = append(answers, ans)
	} else {
		for i := start; i <= end; i++ {
			swap(options, start, i)
			subAnswers := permute(options, start+1)
			for _, ans := range subAnswers {
				answers = append(answers, ans)
			}
			swap(options, start, i)
		}
	}
	return answers
}