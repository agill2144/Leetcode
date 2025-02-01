// tc = o(mn)
// sc = o(mn)
func maxAreaOfIsland(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{
        {-1,0},{1,0},
        {0,-1},{0,1},
    }
    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] != 1 {return 0}

        // logic
        count := 1
        grid[r][c] = 0
        for _, dir := range dirs {
            count += dfs(r+dir[0], c+dir[1])
        }
        return count
    }
    res := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                res = max(res,dfs(i, j))
            }
        }
    }
    return res
}