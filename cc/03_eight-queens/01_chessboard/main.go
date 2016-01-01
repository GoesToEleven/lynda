package main
import "fmt"

// size of board: 8 X 8 in this example
var n = 8

// slice of []int
// slice of slice of int
// position on board: row, column
// eg, board[0][0] means row 0, column 0, eg, the top left cell
// eg, board[8][8] means row 8, column 8, eg, the bottom right cell
// var means zero value, so all values are set to 0
var board = make([][]int, n)

// init does initialization setup
// https://golang.org/doc/effective_go.html#init
// make each inner slice of int
// https://golang.org/doc/effective_go.html#allocation_make
func init() {
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}
}

func main(){

	fmt.Println("The entire date structure")
	fmt.Println(board)

	// r = row, c = column on chessboard
	fmt.Println("Each entry on the chessboard")
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			fmt.Println(r, c, "-", board[r][c])
		}
	}

	fmt.Println("The chessboard")
	for _, v := range board {
		fmt.Println(v)
	}

}