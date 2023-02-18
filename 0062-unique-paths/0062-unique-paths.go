/*
    approach: top down with memoization
    time: o(mn)
    space: o(mn) for memo + m+n for the recursive stack
*/
func uniquePaths(m int, n int) int {
    dr := m-1
    dc := n-1
    memo := make([][]int, m+1)
    for i := 0; i < m+1; i++ { memo[i] = make([]int, n+1)}
    
    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        // base
        if r == dr && c == dc {return 1}
        if r == m || c == n {return 0}

        // logic
        if memo[r][c+1] == 0 { memo[r][c+1] = dfs(r,c+1) }
        if memo[r+1][c] == 0 { memo[r+1][c] = dfs(r+1,c) }
        return memo[r+1][c] + memo[r][c+1]
    }
    return dfs(0,0)
}


/*
    approach: bottom up dp with space optimzation
    - transpose the matrix such that 0,0 is the exit point and solve bottom up
    
    time: o(mn)
    space: o(n)
*/
// func uniquePaths(m int, n int) int {
//     dp := make([]int, n)
    
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if i == 0 {dp[j] = 1; continue}
//             top := dp[j]
//             left := 0
//             if j != 0 {
//                 left = dp[j-1]
//             }
//             dp[j] = top+left
//         }        
//     }
//     return dp[n-1]
    
// }

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