[]int{7, 3, 5, 2, 6, 1, 0, 4}

for row1, col1 := range board {
    for row2 := row1 + 1; row2 < len(board); row2++ {
        col2 := board[row2]
        if intAbs(row2-row1) == intAbs(col2-col1) {
	

ITERATION 1
row1 = 0, col1 = 7 

ITERATION 2
row1 = 1, col1 = 3

ITERATION 3
row1 = 2, col1 = 5

ITERATION 4
row1 = 3, col1 = 2

ITERATION 5
row1 = 4, col1 = 6

ITERATION 6
row1 = 5, col1 = 1

ITERATION 7
row1 = 6, col1 = 0

ITERATION 8
row1 = 7, col1 = 4
