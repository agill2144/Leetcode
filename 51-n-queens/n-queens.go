func solveNQueens(n int) [][]string {
    board := make([][]bool, n)
    for i := 0; i < n; i++ {
        board[i] = make([]bool, n)
    }
    out := [][]string{}
    var dfs func(r int)
    dfs = func(r int) {
        // base
        if r == n {
            tmp := []string{}
            for i := 0; i < len(board); i++ {
                row := new(strings.Builder)
                for j := 0; j < len(board[0]); j++ {
                    if board[i][j] {
                        row.WriteString("Q")
                    } else {
                        row.WriteString(".")
                    }
                }
                tmp = append(tmp, row.String())
                row.Reset()
            }
            out = append(out, tmp)
            return
        }


        // logic
        for j := 0; j < len(board); j++ {
            if isSafe(board, r, j) {
                board[r][j] = true
                dfs(r+1)
                board[r][j] = false
            }
        }
    }
    dfs(0)
    return out
}

func isSafe(board [][]bool, r, c int)bool {
    dirs := [][]int{
        {-1,0},
        {-1,-1},
        {-1,1},
    }
    n := len(board)
    for _, dir := range dirs {
        nr := r + dir[0]
        nc := c + dir[1]
        for nr >= 0 && nc >= 0 && nc < n {
            if board[nr][nc] {return false}
            nr += dir[0]
            nc += dir[1]
        }
    }
    return true
}