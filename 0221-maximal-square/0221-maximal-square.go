
func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {dp[i] = make([]int, n)}
    max := 0    
    dirs := [][]int{{-1,0},{-1,-1},{0,-1}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 {
                dp[i][j] = int(matrix[i][j]-'0')
                if dp[i][j] > max {max = dp[i][j]}
                continue
            }

            if matrix[i][j] == '1' {
                minVal := math.MaxInt64
                for _, dir := range dirs {
                    r := i+dir[0]
                    c := j+dir[1]
                    if r >= 0 && c >= 0 {
                        if dp[r][c] < minVal {minVal = dp[r][c]}
                    } else {
                        minVal = 0
                    }
                }
                dp[i][j] = 1+minVal
                if dp[i][j] > max {max = dp[i][j]}
            }
        }
    }
    return max*max
}