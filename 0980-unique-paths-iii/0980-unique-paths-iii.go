func uniquePathsIII(grid [][]int) int {
    dirs := [][]int{
        {1,0},
        {-1,0},
        {0,1},
        {0,-1},
    }
    sr, sc := -1,-1
    nonObs := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] >= 0 {nonObs++}
            if grid[i][j] == 1 {sr, sc = i,j}
        }
    }
    
    count := 0
    var dfs func(r, c, remain int)
    dfs = func(r, c, remain int) {
        // base
        if r < 0 || r == len(grid) || c < 0 || c == len(grid[0]) || grid[r][c] == -1 || grid[r][c] == 50 {
            return
        }
        if grid[r][c] == 2 && remain == 1 {count++; return}
        
        // logic
        tmp := grid[r][c]
        remain--
        grid[r][c] = 50
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1], remain)
        }
        grid[r][c] = tmp
    }
    dfs(sr, sc, nonObs)
    return count
}