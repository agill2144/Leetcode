// tc = n!
// sc = o(n)
func totalNQueens(n int) int {
    matrix := make([][]bool, n)
    for i := 0; i < n; i++ {matrix[i] = make([]bool,n)}
    
    count := 0
    var dfs func(r int) 
    dfs = func(r int) {
        // base
        if r == n {
            count++
            return
        }

        // logic
        for i := 0; i < n; i++ {
            if isSafe(matrix, r, i) {
                matrix[r][i] = true
                dfs(r+1)
                matrix[r][i] = false
            }
        }
    }
    dfs(0)
    return count
}

func isSafe(matrix [][]bool, r, c int) bool {
    dirs := [][]int{{-1,-1},{-1,1},{-1,0}}
    n := len(matrix)
    for _, dir := range dirs {
        nr := r+dir[0]
        nc := c+dir[1]
        for nr >= 0 && nc >= 0 && nc < n {
            if matrix[nr][nc] {return false}
            nr += dir[0]
            nc += dir[1]
        }
    }
    return true
}