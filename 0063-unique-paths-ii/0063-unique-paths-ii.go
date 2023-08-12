func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {dp[i] = make([]int, n)}
    if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1 {return 0}
    
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {

            if i == m-1 && j == n-1 {dp[i][j] = 1; continue}

            // obstacle cell, skip
            if obstacleGrid[i][j] == 1 {continue}
            
            
            bottom := 0       
            if i+1 < m && obstacleGrid[i+1][j] != 1 {
                bottom = dp[i+1][j]
            }
            right := 0
            if j+1 < n && obstacleGrid[i][j+1] != 1 {
                right = dp[i][j+1]
            }
            dp[i][j] = bottom + right
        }
    }
    return dp[0][0]
}