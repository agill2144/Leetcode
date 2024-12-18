
func change(amount int, coins []int) int {
    m := len(coins)+1
    n := amount+1
    dp := make([]int, n)
    dp[0] = 1
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            coin := coins[i-1]
            am := j
            if coin <= am {
                dp[j] = dp[j] + dp[j-coin] 
            }
        }
    }
    return dp[n-1]
}