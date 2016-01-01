package main

import "fmt"

var n = 8
var board= make([][]int, n)

func init() {
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}
}

func main(){

	fmt.Println("The chessboard")
	for _, v := range board {
		printBoard(v)
		fmt.Print("\n\n\n")
	}

}

func printBoard(board []int) {
	// r = row
	// 16 times b/c 8 rows with data, 8 separator rows
	for r := 0; r < 15; r++ {
		// https://play.golang.org/p/9c79-SGNnH
		if r%2 == 1 {
			fmt.Println("-+-+-+-+-+-+-+-")
		} else {
			// c = column
			// 16 times b/c 8 columns with data, 8 separator columns
			for c := 0; c < 15; c++ {
				if c%2 == 1 {
					fmt.Print("|")
				} else {
					// https://play.golang.org/p/4h2Dp09ul2
					if board[r/2] == c/2 {
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