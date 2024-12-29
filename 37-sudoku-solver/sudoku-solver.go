func solveSudoku(board [][]byte)  {
    n := 9
    rows := make([][]bool, 10)
    cols := make([][]bool, 10)
    boxes := map[string][]bool{}
    for i := 0; i < len(rows); i++ {
        rows[i] = make([]bool, 10)
        cols[i] = make([]bool, 10)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            key := fmt.Sprintf("%v-%v",i/3,j/3)
            if boxes[key] == nil {boxes[key] = make([]bool, 10)}
            if board[i][j] != '.' {
                intVal := int(board[i][j]-'0')
                rows[i][intVal] = true
                cols[j][intVal] = true
                boxes[key][intVal] = true
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
            key := fmt.Sprintf("%v-%v",r/3,c/3)
            inRow := rows[r][i]
            inCol := cols[c][i]
            inBox := boxes[key][i]
            canUse := !inRow && !inCol && !inBox
            if canUse {
                rows[r][i] = true
                cols[c][i] = true
                boxes[key][i] = true
                board[r][c] = byte(i+'0')

                if dfs(r,c+1) {return true}

                board[r][c] = '.'
                boxes[key][i] = false
                cols[c][i] = false
                rows[r][i] = false
            }
        }
        return false
    }
    dfs(0,0)
    
}