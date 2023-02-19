func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1 {return 0}
    
    dp := make([][]int, m+1)
    for i := 0; i < len(dp); i++ {dp[i] = make([]int, n+1)}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if obstacleGrid[i][j] == 1 {dp[i+1][j+1] = -1}
        }
    }
    dp[1][1] = 1
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if i == 1 && j == 1 {continue}
            if dp[i][j] == -1 {continue}
            top := 0
            if dp[i-1][j] != -1 {top = dp[i-1][j]}
            left := 0
            if dp[i][j-1] != -1 {left = dp[i][j-1]}
            dp[i][j] = top+left
        }
    }
    return dp[m][n]
}

// brute force: tle
// func uniquePathsWithObstacles(obstacleGrid [][]int) int {
//     m := len(obstacleGrid)
//     n := len(obstacleGrid[0])
//     var dfs func(r, c int) int
//     dfs = func(r, c int) int {
//         // base
//         if r == m || c == n || obstacleGrid[r][c] == 1 {return 0}
//         // logic
//         if r == m-1 && c == n-1 {
//             return 1
//         }
//         down := dfs(r+1, c)
//         right := dfs(r, c+1)
//         return down + right
//     }
//     return dfs(0,0)
// }