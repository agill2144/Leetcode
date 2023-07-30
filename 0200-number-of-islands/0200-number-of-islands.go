func numIslands(grid [][]byte) int {
    m := len(grid)
    n := len(grid[0])
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] == '0' || visited[r][c] {return}
        
        // logic
        visited[r][c] = true
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' && !visited[i][j] {
                dfs(i, j)
                count++
            }
        }
    }
    return count
}