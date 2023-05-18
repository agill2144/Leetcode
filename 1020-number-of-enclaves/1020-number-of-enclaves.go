func numEnclaves(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] != 1 {return}
        
        // logic
        grid[r][c] = -1
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    
    for i := 0; i < m ; i++ {
        if grid[i][0] == 1 {dfs(i, 0)}
        if grid[i][n-1] == 1 {dfs(i, n-1)}
    }
    
    for j := 0; j < n; j++ {
        if grid[0][j] == 1 {dfs(0, j)}
        if grid[m-1][j] == 1 {dfs(m-1, j)}
    }
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {count++}
        }
    } 
    return count
}