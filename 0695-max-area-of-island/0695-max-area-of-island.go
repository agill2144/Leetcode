func maxAreaOfIsland(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] != 1 {return 0}
        
        // logic
        grid[r][c] = 0 // visited
        count := 1
        for _, dir := range dirs {
            nr := r+dir[0]
            nc := c+dir[1]
            count += dfs(nr, nc)
        }
        return count
    }
    
    ans := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                ans = max(ans, dfs(i, j))    
            }
        }
    }
    return ans
}

func max(x, y int) int {
    if x > y {return x}
    return y
}