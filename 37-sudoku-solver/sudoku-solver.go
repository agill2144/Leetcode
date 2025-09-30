func solveSudoku(board [][]byte)  {
    rows := [10][10]bool{}
    cols := [10][10]bool{}
    boxes := [3][3][10]bool{}
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] != '.' {
                val := int(board[i][j]-'0')
                rows[i][val] = true
                cols[j][val] = true
                boxes[i/3][j/3][val] = true
            }
        }
    }
    n := len(board)
    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        // base
        if c == n {return dfs(r+1, 0)}
        if r == n {return true}
        if board[r][c] != '.' {return dfs(r,c+1)}

        // logic
        for i := 1; i <= 9; i++ {
            inRow := rows[r][i]
            inCol := cols[c][i]
            inBox := boxes[r/3][c/3][i]
            canUse := !inRow && !inCol && !inBox
            if canUse {
                board[r][c] = byte(i+'0')
                rows[r][i] = true
                cols[c][i] = true
                boxes[r/3][c/3][i] = true
                if dfs(r,c+1) {return true}
                board[r][c] = '.'
                rows[r][i] = false
                cols[c][i] = false
                boxes[r/3][c/3][i] = false
            }
        }
        return false
    }
    dfs(0,0)
}