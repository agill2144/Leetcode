func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    m := len(grid1)
    n := len(grid1[0])
    
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    
    // runs over grid2 and checks if each land position 
    // exists as a land in grid1
    // if at any point, thats not true, we turn a flag for this
    // connected component false
    isSubIsland := true
    var dfs func(r, c int)
    dfs = func(r, c int)  {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid2[r][c] != 1 {return}
        
        // logic
        // mark this cell visited
        grid2[r][c] = 0
        
        // check if this position exists as a land in grid1
        if grid1[r][c] != 1 {
            isSubIsland = false
            // but keep going and mark the rest of the connected component visited
        }
        
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    
    }
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid2[i][j] == 1 {
                dfs(i, j)
                if isSubIsland {
                    count++
                }
                // reset
                isSubIsland = true
            }
        }
    }
    return count
}