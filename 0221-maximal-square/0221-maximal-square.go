func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([]int, n+1)
    maxSize := 0
    diagUp := 0
    for i := 1; i < m+1; i++ {
        for j := 1; j < len(dp); j++{
            nextDiagUp := dp[j]
            if matrix[i-1][j-1] == '1' {
                dp[j] = min(diagUp, min(dp[j], dp[j-1]))+1
                if dp[j] > maxSize {maxSize = dp[j]}
            } else {
                dp[j] = 0
            }
            diagUp = nextDiagUp
        }
    }
    return maxSize*maxSize
}

func min(x, y int) int {
    if x < y {return x}
    return y
}