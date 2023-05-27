func longestIncreasingPath(matrix [][]int) int {
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    m := len(matrix)
    n := len(matrix[0])
    maxCount := 0
    
    memo := make([][]int, m)
    for i := 0; i < m; i++ {memo[i] = make([]int, n)}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            memo[i][j] = -1 // not-yet-solved
        }
    }
    
    var dfs func(r, c, count int) int
    dfs = func(r, c, count int) int {
        // base
        
        // logic
        if memo[r][c] != -1 {
            return memo[r][c]
        } else {
            ans := count
            for _, dir := range dirs {
                nr := r+dir[0]
                nc := c+dir[1]
                if nr >= 0 && nr < m && nc >= 0 && nc < n && matrix[nr][nc] > matrix[r][c] {        
                    ans = max(ans, dfs(nr, nc, 1)+1)
                }
            }
            memo[r][c] = ans
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