package sudoku

type SudokuOverlap struct {
	boards [3]Sudoku
}

// Returns a sudoku board with all tiles empty and
func NewSudokuOverlap(board1 Sudoku, board2 Sudoku, board3 Sudoku) SudokuOverlap {
	return SudokuOverlap{
		boards: [3]Sudoku{board1, board2, board3},
	}
}