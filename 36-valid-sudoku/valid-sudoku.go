func isValidSudoku(board [][]byte) bool {
    n := len(board)
    boxSet := map[string][]bool{}
    for i := 0; i < n; i++ {
        rowSet := make([]bool, n+1)
        colSet := make([]bool, n+1)
        for j := 0; j < n; j++ {
            if board[i][j] != '.'{
                rowVal := int(board[i][j]-'0')
                if rowSet[rowVal] {return false}
                rowSet[rowVal] = true

                boxKey := fmt.Sprint("%v-%v", i/3,j/3)
                if boxSet[boxKey] == nil {boxSet[boxKey] = make([]bool, n+1)}
                if boxSet[boxKey][rowVal] {return false}
                boxSet[boxKey][rowVal] = true

            }

            if board[j][i] != '.' {
                colVal := int(board[j][i]-'0')
                if colSet[colVal] {return false}
                colSet[colVal] = true
            }
        }
    }
    return true
}
