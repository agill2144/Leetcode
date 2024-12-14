func solveSudoku(board [][]byte)  {
    rows := [10][10]bool{}
    cols := [10][10]bool{}
    box :=  [3][3][10]bool{}
    n := len(board)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] != '.' {
                idx := int(board[i][j]-'0')
                rows[i][idx] = true
                box[i/3][j/3][idx] = true
            }
            if board[j][i] != '.' {
                idx := int(board[j][i]-'0')
                cols[i][idx] = true
            }
        }
    }
    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        // base
        if r == n {return true}
        if c == n {return dfs(r+1, 0)}
        if board[r][c] != '.' {return dfs(r,c+1)}

        // logic
        for i := 1; i <= 9; i++ {
            inRow := rows[r][i]
            inCol := cols[c][i]
            inBox := box[r/3][c/3][i]
            canUse := !inRow && !inCol && !inBox

            if canUse {
                rows[r][i] = true
                cols[c][i] = true
                box[r/3][c/3][i] = true
                board[r][c] =  byte(i)+'0'
                if dfs(r, c+1) {return true}
                board[r][c] = '.'
                box[r/3][c/3][i] = false
                cols[c][i] = false
                rows[r][i] = false
            }
        }
        return false
    }
    dfs(0, 0)
}