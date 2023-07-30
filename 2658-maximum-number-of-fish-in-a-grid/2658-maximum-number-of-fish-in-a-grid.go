func findMaxFish(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    
    maxFishes := 0
    runningTotal := 0
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] == 0 {return}
        
        
        // logic
        // add to runningTotal and mark it visited
        runningTotal += grid[r][c]
        if runningTotal > maxFishes {maxFishes = runningTotal}
        grid[r][c] = 0
        
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
        
    }
    
    for i := 0; i < m ; i++ {
        for j := 0; j < n ; j++ {
            if grid[i][j] != 0 {
                runningTotal = 0
                dfs(i, j)
            }
        }
    }
    return maxFishes
}