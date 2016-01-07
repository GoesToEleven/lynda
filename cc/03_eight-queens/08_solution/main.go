package main

import "fmt"

func main() {
	// stores column which queen is in for each row
	// index is row, column is value
	options := []int{0, 1, 2, 3, 4, 5, 6, 7}
	perm := getOptions(options)
	var numAns int
	for _, v := range perm {
		numAns++
		printBoard(v)
		fmt.Print("\n\n\n")
	}
	fmt.Println("Number of answers:", numAns)
}

func printBoard(board []int) {
	for y := 0; y < 15; y++ {
		if y%2 == 1 {
			fmt.Println("-+-+-+-+-+-+-+-")
		} else {
			for x := 0; x < 15; x++ {
				if x%2 == 1 {
					fmt.Print("|")
				} else {
					if board[y/2] == x/2 {
						fmt.Print("Q")
					} else {
						fmt.Print(" ")
					}
				}
			}
			fmt.Println()
		}
	}
}

func getOptions(options []int) [][]int {
	perm := permutations(options)
	ret := [][]int{}
	for _, option := range perm {
		if checkValid(option) {
			ret = append(ret, option)
		}
	}
	return ret
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkValid(board []int) bool {
	for row1, col1 := range board {
		for row2 := row1 + 1; row2 < len(board); row2++ {
			col2 := board[row2]
			if intAbs(row2-row1) == intAbs(col2-col1) {
				return false
			}
		}
	}
	return true
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
