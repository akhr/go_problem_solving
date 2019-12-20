package validsudoku36

func isValidSudoku(board [][]byte) bool {
	return areColumnsValid(board) && areRowsValid(board) && areSubBlocksValid(board)
}

func areColumnsValid(board [][]byte) bool {
	for c := 0; c < 9; c++ {
		set := make([]int, 10)
		for i := 0; i < 9; i++ {
			if !isValidCell(set, board[i][c]) {
				return false
			}
		}
	}
	return true
}

func areRowsValid(board [][]byte) bool {
	for r := 0; r < 9; r++ {
		set := make([]int, 10)
		for i := 0; i < 9; i++ {
			if !isValidCell(set, board[r][i]) {
				return false
			}
		}
	}
	return true
}

func areSubBlocksValid(board [][]byte) bool {
	for r := 0; r < 9; r = r + 3 {
		for c := 0; c < 9; c = c + 3 {

			set := make([]int, 10)
			for i := r; i < (r + 3); i++ {
				for j := c; j < (c + 3); j++ {
					if !isValidCell(set, board[i][j]) {
						return false
					}
				}
			}
		}
	}
	return true
}

func isValidCell(set []int, cellVal byte) bool {
	if cellVal == '.' {
		return true
	}
	if set[cellVal-48] == 1 { //ASCII of '0' = 48. Subtract 48 to get actual int val
		return false
	}
	set[cellVal-48] = 1
	return true
}
