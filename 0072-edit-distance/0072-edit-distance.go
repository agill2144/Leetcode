func minDistance(word1 string, word2 string) int {
    m := len(word1)+1
    n := len(word2)+1
    dp := make([]int, n)
    for i := 0; i < n ; i++ {dp[i] = i}

    for i := 1; i < m; i++ {
        // new row, since j loop starts from 1
        // we know that diagUp is i-1;j-1
        // in 1d array thats dp[0]
        // save that FIRST
        diagUp := dp[0]
        // why becuase, we need the updated left value
        // dp[0] in each new row, will be whatever i was ( look at dp matrix )
        // so we need old dp[0] as a diagUp 
        // and we need new dp[0] for the left value
        dp[0] = i
        for j := 1; j < n; j++ {
            nextDiagUp := dp[j]
            if word1[i-1] == word2[j-1] {
                dp[j] = diagUp
            } else {
                dp[j] = min(diagUp, min(dp[j], dp[j-1]))+1
            }
            diagUp = nextDiagUp
        }
    }
    return dp[n-1]
}

func min(x, y int) int {
    if x < y {return x}
    return y
}