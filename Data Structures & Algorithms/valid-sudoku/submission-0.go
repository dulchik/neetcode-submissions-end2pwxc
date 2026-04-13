func isValidSudoku(board [][]byte) bool {
    rows := make([][]bool, 9)
    cols := make([][]bool, 9)
    boxes := make([][]bool, 9)

    for i := 0; i < 9; i++ {
        rows[i] = make([]bool, 9)
        cols[i] = make([]bool, 9)
        boxes[i] = make([]bool, 9)
    }

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] != '.' {
                num := board[i][j] - '1'
                boxIndex := (i/3)*3 + (j/3)
                if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
                    return false
                }
                rows[i][num] = true
                cols[j][num] = true
                boxes[boxIndex][num] = true
            }
        }
    }
    return true
}