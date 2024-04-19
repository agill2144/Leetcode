func numIslands(grid [][]byte) int {
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(r, c int) 
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] != '1' {
            return
        }

        // logic
        // mark it visited
        grid[r][c] = '0'
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                dfs(i, j)
                count++
            }
        }
    }
    return count
}