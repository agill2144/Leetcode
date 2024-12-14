func solveSudoku(board [][]byte)  {
    rows := make([][]bool, 10)
    cols := make([][]bool, 10)
    box := map[string][]bool{}
    for i := 0; i < len(rows); i++ {
        rows[i] = make([]bool, 10)
        cols[i] = make([]bool, 10)
    }
    n := len(board)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            boxKey := fmt.Sprintf("%v-%v", i/3, j/3)
            if box[boxKey] == nil {box[boxKey] = make([]bool, 10)}
            if board[i][j] != '.' {
                idx := int(board[i][j]-'0')
                rows[i][idx] = true
                box[boxKey][idx] = true
            }
            if board[j][i] != '.' {
                idx := int(board[j][i]-'0')
                cols[i][idx] = true
            }
        }
    }
    choices := []byte{'1','2','3','4','5','6','7','8','9'}
    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        // base
        if r == n {return true}
        if c == n {return dfs(r+1, 0)}
        if board[r][c] != '.' {return dfs(r,c+1)}

        // logic
        for i := 0; i < len(choices); i++ {
            choice := choices[i]
            idx := int(choice-'0')
            boxKey := fmt.Sprintf("%v-%v", r/3, c/3)
            if box[boxKey] == nil {box[boxKey] = make([]bool, 10)}
            
            inRow := rows[r][idx]
            inCol := cols[c][idx]
            inBox := box[boxKey][idx]
            canUse := !inRow && !inCol && !inBox

            if canUse {
                rows[r][idx] = true
                cols[c][idx] = true
                box[boxKey][idx] = true
                board[r][c] = choice
                if dfs(r, c+1) {return true}
                board[r][c] = '.'
                box[boxKey][idx] = false
                cols[c][idx] = false
                rows[r][idx] = false
            }
        }
        return false
    }
    dfs(0, 0)
}