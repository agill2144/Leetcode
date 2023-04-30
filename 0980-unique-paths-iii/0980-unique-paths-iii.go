// "walk over every non-obstacle square exactly once" in a path
// Note: WE MUST VISIT ALL CELLS THAT ARE REACHABLE FOR A PATH TO COUNT AS 1
func uniquePathsIII(grid [][]int) int {
    
    dirs := [][]int{
        {1,0},
        {-1,0},
        {0,1},
        {0,-1},
    }
    sr, sc := -1,-1
    reachableCells := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] >= 0 {reachableCells++}
            if grid[i][j] == 1 {sr, sc = i,j}
        }
    }
    
    count := 0
    var dfs func(r, c, cellsReached int)
    dfs = func(r, c, cellsReached int) {
        // base
        if r < 0 || r == len(grid) || c < 0 || c == len(grid[0]) || grid[r][c] == -1 || grid[r][c] == 50 {
            return
        }
        if grid[r][c] == 2 && cellsReached == reachableCells {count++; return}
        
        // logic
        tmp := grid[r][c]
        cellsReached++
        grid[r][c] = 50
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1], cellsReached)
        }
        grid[r][c] = tmp
    }
    dfs(sr, sc, 1)
    return count
}