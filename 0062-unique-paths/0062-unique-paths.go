func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := 0; i < len(dp); i++ {dp[i] = make([]int, n)}
    dp[0][0] = 1
    
    for i := 0; i < len(dp); i++ {
        for j := 0; j < len(dp[0]); j++ {
            if i == 0 && j == 0 {continue}
            if i == 0 {dp[i][j] = 1; continue}
            if j == 0 {dp[i][j] = 1; continue}
            dp[i][j] = dp[i][j-1]+dp[i-1][j]
        }
    }
    return dp[m-1][n-1]
    
}