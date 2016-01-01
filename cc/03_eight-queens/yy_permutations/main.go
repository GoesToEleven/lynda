package main

import "fmt"

func main() {
	options := "ABC"
	n := len(options)
	// takes string and starting index:
	permute(options, 0)

	perm := getOptions(options)
	var numAns int
	for _, v := range perm {
		numAns++
		printBoard(v)
		fmt.Print("\n\n\n")
	}
	fmt.Println("Number of answers:", numAns)
}

func permute(str string, start int) []byte {
	bs := []byte(str)
	answers := []byte
	end := len(bs) - 1
	if start == end {
		answers = append(answers, bs)
	} else {
		for i := start; i <= end; i++ {
			swap(bs, start, i)
			fmt.Println("FIRST SWAP", i, bs) // DEBUGGING
			subAnswers := permute(bs, start+1)
			fmt.Println(subAnswers) // DEBUGGING
			for _, ans := range subAnswers {
				answers = append(answers, ans)
			}
			swap(bs, start, i)
			fmt.Println("SECOND SWAP", i, bs) // DEBUGGING
		}
	}
	return answers
}

func swap(bs []byte, x, y int) {
	bs[x], bs[y] = bs[y], bs[x]
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
		for row2 := row1 + 1; row2 < 8; row2++ {
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



/*
credit for algorithm:
http://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string/
*/
