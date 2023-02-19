/*
    approach: bottom up dp with space optimization
    - we just need access to previous row for each row we solve
    - therefore we do not need the full mxn dp matrix
    - we can do the same with 1d array
    - before we mutate;
        - the current 1d array is THE PREVIOUS / ABOVE ROW IN DP MATRIX
        - therefore top = dp[j]
        - and left = dp[j-1]
    
    time: o(mn)
    space: o(n)
*/
func uniquePaths(m int, n int) int {
    dp := make([]int, n+1)
    dp[1] = 1
    for i := 0; i < m; i++ {
        for j := 1; j < len(dp); j++ {
            if i == 1 && j == 1 {continue}
            left := dp[j-1]
            top := dp[j]
            dp[j] = left+top
        }
    }
    return dp[n]
}


/*
    approach: bottom up DP
    - there are repeating subproblems.
    - solve the smallest problem and use its answer to solve bigger and bigger problems iteratively ( start small -> big )
    - how, if we have already solved for a cell , we can just use the value of that cell and not compute from that cell further.
    - transpose the matrix such that we start from 0,0
    - if 0,0 is our destination, there is only 1 way to reach it.
        - therefore value at 0,0 will be 1
    - then loop over the rest of the matrix, solving each cells problem by looking back
        - question to ask at each cell, how many ways we can reach this cell if my subproblem cell are on top and left ?
        - ans: top+left
        - repeat for the rest of the cells from 0,0
        - tip: only look back at the current problem from current i,j and forget about the rest of the right side
    - matrix DP observation tricks;
        - may involve 2 passes from 0,0 and from m-1, n-1 if looking for min/max something that depends on neighboring cells.
            - example: 0/1 matrix
        - if starting from 0,0, to look back; top and left
        - if start from m-1,n-1, to look back; down and right
        - this is not looking for min/max so no 2 passes needed from 0,0, m-1,n-1
    
    time: o(mn)
    space: o(mn)
*/
// func uniquePaths(m int, n int) int {
//     dp := make([][]int, m+1)
//     for i := 0; i < m+1; i++ {dp[i] = make([]int, n+1)}
    
//     dp[1][1] = 1
//     for i := 1; i < len(dp); i++ {
//         for j := 1; j < len(dp[i]); j++ {
//             if i == 1 && j == 1 {continue}
//             dp[i][j] = dp[i-1][j] + dp[i][j-1]
//         }
//     }
//     return dp[m][n]
// }

/*
    approach: brute force dfs
    - explore all paths recursively
    - each cell has 2 option ; go down or go right
    - each cell will explore all of those 2 options
    - there are mxn cells
    
    time: 2^mn
    space: o(max(m,n)); draw a wide matrix and a narrow matrix and track stack, it helps prove this 
*/
// func uniquePaths(m int, n int) int {
//     var dfs func(r, c int) int
//     dfs = func(r, c int) int {
//         // base
//         if r == m-1 && c == n-1 {return 1}
//         if r == m || c == n {return 0}
        
//         // logic
//         return dfs(r+1,c) + dfs(r, c+1)
//     }
//     return dfs(0, 0)
// }