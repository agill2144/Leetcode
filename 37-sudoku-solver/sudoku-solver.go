func solveSudoku(board [][]byte)  {
    n := 9
    rows := [10][10]bool{}
    cols := [10][10]bool{}
    boxes := [3][3][10]bool{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] != '.' {
                intVal := int(board[i][j]-'0')
                rows[i][intVal] = true
                cols[j][intVal] = true
                boxes[i/3][j/3][intVal] = true
            }
        }
    }
    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        // base
        if r == n-1 && c == n {return true}
        if c == n {
            return dfs(r+1, 0)
        }
        if board[r][c] != '.' {
            return dfs(r, c+1)
        }
        

        // logic
        for i := 1; i <= 9; i++ {
            inRow := rows[r][i]
            inCol := cols[c][i]
            inBox := boxes[r/3][c/3][i]
            canUse := !inRow && !inCol && !inBox
            if canUse {
                rows[r][i] = true
                cols[c][i] = true
                boxes[r/3][c/3][i] = true
                board[r][c] = byte(i+'0')

                if dfs(r,c+1) {return true}

                board[r][c] = '.'
                boxes[r/3][c/3][i] = false
                cols[c][i] = false
                rows[r][i] = false
            }
        }
        return false
    }
    dfs(0,0)
    
}