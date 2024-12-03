func solveSudoku(board [][]byte)  {
    row := map[int]map[byte]bool{}
    col := map[int]map[byte]bool{}    
    box := map[string]map[byte]bool{}
    m := len(board)
    n := len(board[0])

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if row[i] == nil {row[i] = map[byte]bool{}}
            if col[i] == nil {col[i] = map[byte]bool{}}
            boxKey := fmt.Sprintf("%v-%v",i/3, j/3)
            if box[boxKey] == nil {box[boxKey] = map[byte]bool{}}
            rowVal := board[i][j]
            if rowVal != '.' {row[i][rowVal] = true; box[boxKey][rowVal] = true}
            colVal := board[j][i]
            if colVal != '.' {col[i][colVal] = true}
        }
    }
    choices := []byte{'1','2','3','4','5','6','7','8','9'}
    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        // base
        if r == m {return true}
        if c == n {return dfs(r+1, 0)}

        if board[r][c] != '.' {return dfs(r,c+1)}
        // logic
        boxKey := fmt.Sprintf("%v-%v", r/3,c/3)
        rowSet := row[r]
        colSet := col[c]
        boxSet := box[boxKey]
        for i := 0; i < len(choices); i++ {
            choice := choices[i]
            inRow := rowSet[choice]
            inCol := colSet[choice]
            inBox := boxSet[choice]
            canUse := !inRow && !inCol && !inBox
            if canUse {

                // action
                board[r][c] = choice
                rowSet[choice] = true
                colSet[choice] = true
                boxSet[choice] = true

                // recurse
                if dfs(r,c+1) {return true}

                // backtrack
                delete(boxSet, choice)
                delete(colSet, choice)
                delete(rowSet,choice)
                board[r][c] = '.'
            }
        }
        return false
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == '.' {
                if dfs(i, j) {return}
            }
        }
    }
}


