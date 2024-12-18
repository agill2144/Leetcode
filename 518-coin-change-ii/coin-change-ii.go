
func change(amount int, coins []int) int {
    m := len(coins)+1
    n := amount+1
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
        dp[i][0] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            coin := coins[i-1]
            am := j
            if coin > am {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] + dp[i][j-coin] 
            }
        }
    }
    return dp[m-1][n-1]
}