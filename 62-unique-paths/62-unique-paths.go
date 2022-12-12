/*
    There are repeated subproblems in this, therefore we can use DP
    
    approach: bottom up
    - We will start by solving the full problem for each individual cell
    - As if this individual cell is the destination - we will do this for each cell from left to right
    - "How many ways can this i,j cell can be reached"
    - Since bottom up means looking back/left at previous decisions and optimizing your current decision
        - the optimal decision at each cell is just a sum of how many ways this cell can be reached from immediate left and immediate top neighbors.
    - So this means, that for each cell we will look up + look left and this gives us the answer for "How many ways can this i,j cell can be reached"
    - What about the 0th row, which has no row above... and 0th col which can no left neighbors?
        - Thats why we will take an extra row and col - dummy row and dummy col
    - So we will fill out a dp matrix bottom up way 
        - traverse matrix
        - Solve the main question for each individual cell / subproblem by looking left and above ( dp[i-1][j] + dp[i][j-1] )
    - Then finally the last row, last col will have our answer.
    
    time: o(mn)
    space: o(mn)
*/

// func uniquePaths(m int, n int) int {
//     dp := make([][]int, m+1)
//     for idx, _ := range dp {
//         dp[idx] = make([]int, n+1)
//     }
//     dp[1][1] = 1
//     for i := 1; i < len(dp); i++ {
//         for j := 1; j < len(dp[0]); j++ {
//             if i == 1 && j == 1 {continue}
//             dp[i][j] = dp[i-1][j] + dp[i][j-1]
//         }
//     }
//     return dp[len(dp)-1][len(dp[0])-1]
// }

/*
    approach: bottom up dp
    - same thing as above but using 1D array
    time: o(mn)
    space: o(n)
*/
func uniquePaths(m int, n int) int {
    dp := make([]int, n+1)
    dp[1] = 1
    for i := 0; i < m; i++ {
        for j := 1; j < len(dp); j++ {
            if j == 1 {continue}
            dp[j] = dp[j] + dp[j-1]
        }
    }
    return dp[len(dp)-1]
}


// top down
// time and space: o(mn)
// func uniquePaths(m int, n int) int {
//     dp := make([][]int, m+1)
//     for idx, _ := range dp {
//         dp[idx] = make([]int, n+1)
//     }
//     // there is only 1 way to get out from the last cell ( which is right ) therefore only 1 way
//     dp[m-1][n-1] = 1
//     for i := m-1; i >= 0; i-- {
//         for j := n-1; j >= 0; j-- {
//             if i == m-1 && j == n-1 {continue}
//             dp[i][j] = dp[i+1][j] + dp[i][j+1]
//         }
//     }
//     return dp[0][0]
// }