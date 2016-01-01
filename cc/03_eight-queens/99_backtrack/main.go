package main

// size of board: 8 X 8 in this example
var n = 8

// slice of []int
// slice of slice of int
// position on board: row, column
// eg, board[0][0] means row 0, column 0, eg, the top left cell
// eg, board[8][8] means row 8, column 8, eg, the bottom right cell
// var means zero value, so all values are set to 0
var board = make([][]int, n)

func init() {
	for i := 0; i < n; i++ {
		board[i] = make([]int, 1)
	}
}

func main(){

}

func queenSafe(row, col int) bool {
	var i int
	for i <= n {
		if board[row][i] != 0 || board[i][col] != 0 {

		}
	}
}

