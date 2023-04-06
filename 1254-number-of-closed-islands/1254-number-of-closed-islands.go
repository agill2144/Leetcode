func closedIsland(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    
    // mark all edge connected components
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] != 0 {
            return
        }
        
        // logic
        grid[r][c] = -1
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if (i == 0 || i == m-1 || j == 0 || j == n-1) && grid[i][j] == 0 {
                dfs(i, j)
            }
        }
    }
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                count++
                dfs(i, j)
            }
        }
    }
    
    return count
}