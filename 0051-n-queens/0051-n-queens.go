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

            for i := 0; i < n; i++ {
                rStr := new(strings.Builder)
                for j := 0; j < n; j++ {
                    if board[i][j] {
                        rStr.WriteString("Q")
                    } else {
                        rStr.WriteString(".")
                    }
                }
                tmp = append(tmp, rStr.String())
            }
            out = append(out, tmp)
            return
        }
        
        // logic
        for c := 0; c < n; c++ {
            if isSafe(r,c,board) {
                board[r][c] = true
                dfs(r+1)
                board[r][c] = false
            }
        }
    }
    dfs(0)
    return out
}

func isSafe(r, c int, board [][]bool) bool {
    dirs := [][]int{{-1,-1},{-1,0},{-1,1}}
    n := len(board)
    for _, dir := range dirs {
        nr := r + dir[0]
        nc := c + dir[1]
        for nr >= 0 && nc >= 0 && nr < n && nc < n {
            if board[nr][nc] {return false}
            nr += dir[0]
            nc += dir[1]
        }
    }
    return true
}