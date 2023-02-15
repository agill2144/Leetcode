
func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {dp[i] = make([]int, n+1)}
    max := 0 
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if matrix[i-1][j-1] == '1' {
                dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j],dp[i][j-1]))+1
                if dp[i][j] > max {max = dp[i][j]}
            }
        }
    }
    return max*max
}

func min(x, y int) int {
    if x < y {return x}
    return y
}