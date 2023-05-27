func countPaths(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {dp[i] = make([]int, n)}
    upDirs := [][]int{{-1,0},{0,-1}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            curr := grid[i][j]
            total := 0
            for _, dir := range upDirs {
                r := i + dir[0]
                c := j + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] > curr {
                    total++
                }
            }
            dp[i][j] = total+1
        }
    }
    count := 0    
    downDirs := [][]int{{1,0},{0,1}}
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            curr := grid[i][j]
            total := dp[i][j]
            for _, dir := range downDirs {
                r := i + dir[0]
                c := j + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] > curr {
                    total+=dp[r][c]
                }
            }
            dp[i][j] = total
            count += dp[i][j]
        }
    }
    return count
}