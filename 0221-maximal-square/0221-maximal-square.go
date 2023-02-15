func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    maxSize := 0
    dp := make([]int, n+1)
    diagUp := 0    
    for i := 1; i < m+1; i++ {
        for j := 1; j < n+1; j++ {
            newDiagUp := dp[j]
            if matrix[i-1][j-1] == '1' {
                dp[j] = int(math.Min(float64(dp[j-1]), math.Min(float64(dp[j]), float64(diagUp))))+1
            } else {
                dp[j] = 0
            }
            if dp[j] > maxSize {maxSize = dp[j]}
            diagUp = newDiagUp
        }
    }
    return maxSize*maxSize
}