package main

import "fmt"

func main() {
	options := []string{"A", "B", "C"}
	perm := permutations(options)
	for _, v := range perm {
		fmt.Println(v)
	}
	fmt.Println("Number of answers:", len(perm))
}

func permutations(options []string) [][]string {
	return permute(options, 0)
}

func swap(a []string, x, y int) {
	a[x], a[y] = a[y], a[x]
}

func permute(options []string, start int) [][]string {
	answers := [][]string{}
	end := len(options) - 1
	if start == end {
		ans := make([]string, len(options))
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
