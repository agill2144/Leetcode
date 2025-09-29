func solveSudoku(board [][]byte)  {
    n := len(board)
    rows := [10][10]bool{}
    cols := [10][10]bool{}
    boxes := [3][3][10]bool{}
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] != '.' {
                intVal := int(board[i][j]-'0')
                r, c := i/3, j/3
                rows[i][intVal] = true
                cols[j][intVal] = true
                boxes[r][c][intVal] = true
            }
        }
    }
    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        // base
        if c == n {
            return dfs(r+1, 0)
        }
        if r == n {
            return true
        }
        if board[r][c] != '.' {
            return dfs(r, c+1)
        }

        // logic
        for i := 1; i <= 9; i++ {
            inRow := rows[r][i]
            inCol := cols[c][i]
            inBox := boxes[r/3][c/3][i]
            if !inRow && !inCol && !inBox {
                rows[r][i] = true
                cols[c][i] = true
                boxes[r/3][c/3][i] = true
                board[r][c] = byte(i+'0')
                if dfs(r, c+1) {return true}
                rows[r][i] = false
                cols[c][i] = false
                boxes[r/3][c/3][i] = false
                board[r][c] = '.'
            }
        }
        return false
    }
    dfs(0,0)
}