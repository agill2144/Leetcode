func solveNQueens(n int) [][]string {
    matrix := make([][]bool, n)
    for i := 0; i < n; i++ {matrix[i] = make([]bool, n)}
    dirs := [][]int{{-1,0},{-1,-1},{-1,1}}
    out := [][]string{}
    var dfs func(r int)
    dfs = func(r int) {
        // base
        if r == n {
            tmp := []string{}
            for i := 0; i < n; i++ {
                rowStr := new(strings.Builder)
                for j := 0; j < n; j++ {
                    if matrix[i][j]{
                        rowStr.WriteString("Q")
                    } else {
                        rowStr.WriteString(".")
                    }
                }
                tmp = append(tmp, rowStr.String())
            }
            out = append(out, tmp)
            return
        }
        
        // logic
        for j := 0; j < n; j++ {
            if isSafe(r, j, matrix, dirs) {
                matrix[r][j] = true
                dfs(r+1)
                matrix[r][j] = false
            }
        }
    }
    dfs(0)
    return out
}

func isSafe(r, c int, matrix [][]bool, dirs [][]int) bool {
    for _, dir := range dirs {
        nr := r + dir[0]
        nc := c + dir[1]
        for nr >= 0 && nc >= 0 && nc < len(matrix[0]) {
            if matrix[nr][nc] {return false}
            nr += dir[0]
            nc += dir[1]
        }
    }
    return true
}