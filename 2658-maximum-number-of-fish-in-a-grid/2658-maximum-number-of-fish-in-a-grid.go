func findMaxFish(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    
    maxFishes := 0
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] == 0 {return 0}
        
        
        // logic
        // add to runningTotal and mark it visited
        runningTotal := grid[r][c]
        grid[r][c] = 0
        
        for _, dir := range dirs {
            runningTotal += dfs(r+dir[0], c+dir[1])
        }
        return runningTotal
        
    }
    
    for i := 0; i < m ; i++ {
        for j := 0; j < n ; j++ {
            if grid[i][j] != 0 {
                if val := dfs(i, j); val > maxFishes {
                    maxFishes = val
                }
            }
        }
    }
    return maxFishes
}