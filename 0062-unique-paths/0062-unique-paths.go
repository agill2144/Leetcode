/*
    approach: bottom up dp with space optimzation
    - transpose the matrix such that 0,0 is the exit point and solve bottom up
    
    time: o(mn)
    space: o(mn)
*/
func uniquePaths(m int, n int) int {
    dp := make([]int, n)
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 {dp[j] = 1; continue}
            top := dp[j]
            left := 0
            if j != 0 {
                left = dp[j-1]
            }
            dp[j] = top+left
        }        
    }
    return dp[n-1]
    
}

/*
    approach: bottom up dp 
    - transpose the matrix such that 0,0 is the exit point and solve bottom up
    
    time: o(mn)
    space: o(mn)
*/
// func uniquePaths(m int, n int) int {
//     dp := make([][]int, m)
//     for i := 0; i < len(dp); i++ {dp[i] = make([]int, n)}
//     dp[0][0] = 1
    
//     for i := 0; i < len(dp); i++ {
//         for j := 0; j < len(dp[0]); j++ {
//             if i == 0 && j == 0 {continue}
//             if i == 0 {dp[i][j] = 1; continue}
//             if j == 0 {dp[i][j] = 1; continue}
//             dp[i][j] = dp[i][j-1]+dp[i-1][j]
//         }
//     }
//     return dp[m-1][n-1]
    
// }