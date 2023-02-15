func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    maxSize := 0
    dp := make([][]int, m+1)
    for idx, _ := range dp {dp[idx] = make([]int, n+1)}
    
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if matrix[i-1][j-1] == '1' {
                dp[i][j] = int(math.Min(float64(dp[i-1][j-1]), math.Min(float64(dp[i][j-1]), float64(dp[i-1][j]))))+1
                if dp[i][j] > maxSize {
                    maxSize = dp[i][j]
                }
            }
        }
    }
    return maxSize*maxSize
}