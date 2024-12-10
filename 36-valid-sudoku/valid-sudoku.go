func isValidSudoku(board [][]byte) bool {
    n := len(board)
    // tc = o(n^2)
    // sc = o(1)

    // there are only 9 possible keys
    // and each key only has a array of size 10
    // i.e constant space
    // sc = o(1)
    boxSet := map[string][]bool{}
    for i := 0; i < n; i++ {
        // constant space ; o(1)
        rowSet := make([]bool, n+1)
        // constant space ; o(1)
        colSet := make([]bool, n+1)
        for j := 0; j < n; j++ {
            if board[i][j] != '.'{
                rowVal := int(board[i][j]-'0')
                if rowSet[rowVal] {return false}
                rowSet[rowVal] = true

                boxKey := fmt.Sprintf("%v-%v", i/3,j/3)
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
