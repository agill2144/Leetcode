func numEnclaves(grid [][]int) int {
    dirs := [][]int{
        {1,0}, {-1,0}, {0,-1},{0,1},
    }
    m := len(grid)
    n := len(grid[0])
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r == m || r < 0 || c == n || c < 0 || grid[r][c] != 1 {return}
        
        // logic
        grid[r][c] = -1
        for _, dir := range dirs{
            dfs(r+dir[0], c+dir[1])
        }
    }
    for j := 0; j < n; j++ {dfs(0, j); dfs(m-1, j)}
    for i := 0; i < m; i++ {dfs(i, 0); dfs(i, n-1)}
    
    count := 0
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if grid[i][j] == 1 {count++}
        }
    }
    return count
}