func getMaximumGold(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    maxGold := 0

    var dfs func(r, c int, gold int)
    dfs = func(r, c int, gold int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] == 0 {return}

        // logic
        tmp := grid[r][c]
        gold += grid[r][c]
        maxGold = max(maxGold, gold)
        grid[r][c] = 0
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1], gold)
        }
        grid[r][c] = tmp
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != 0 {dfs(i, j, 0)} 
        } 
    }
    return maxGold
} 