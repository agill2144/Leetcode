func longestIncreasingPath(matrix [][]int) int {
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    m := len(matrix)
    n := len(matrix[0])
    maxCount := 0
    
    // all vals will be 0; i.e not-yet-solved
    memo := make([][]int, m)
    for i := 0; i < m; i++ {memo[i] = make([]int, n)}
    

    var dfs func(r, c, count int) int
    dfs = func(r, c, count int) int {
        // base
        // already solved? return the solved value
        if memo[r][c] != 0 {
            return memo[r][c]
        }
        
        // logic
        // otherwise solve for this cell, initially the count will be 1
        memo[r][c] = 1
        for _, dir := range dirs {
            nr := r+dir[0]
            nc := c+dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && matrix[nr][nc] > matrix[r][c] {        
                memo[r][c] = max(memo[r][c], dfs(nr, nc, 1)+1)
            }
        }
        return memo[r][c]
    }
    

    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dfs(i, j, 1)
            if memo[i][j] > maxCount {maxCount = memo[i][j]} 
        }
    }
    return maxCount
}

func max(x, y int) int {
    if x > y {return x}
    return y
}