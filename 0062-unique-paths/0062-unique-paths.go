/*
    approach: bottom up dp with space optimization
    - we just need access to previous row for each row we solve
    - therefore we do not need the full mxn dp matrix
    - we can do the same with 1d array
    - before we mutate;
        - the current 1d array is THE PREVIOUS / ABOVE ROW IN DP MATRIX
        - therefore top = dp[j]
        - and left = dp[j-1]
    - the 0th row will only have 1 way ( check dp matrix )
    - and the 0th col in each row will always have 1 way ( top way )
    - therefore we can prefill 1d dp array with 1s
        - dp array of size n
    - then we can skip row 0, because we just solved row 0 by prefilling 1s because there is only 1 way if our subproblem is the JUST THE 0th row
    - we can also skip col 0 in each row, since it does not have a left choice and only top choice which will always be 1 ( look at dp matrix for ref )
    - dp[j] = dp[j](top) + dp[j-1](left)
    
    can we do with 2 vars ?
    - no, we keep overriding those 2 vars and when we get to next row,
    -  we need those 2 vars to hold the previous row start values
    - and they wont, these 2 ars will hold the last values from the previous row
    - for matrixes, 2 vars have not worked because of this
    - for 1d array input array its possible
    
    time: o(mn)
    space: o(n)
*/
func uniquePaths(m int, n int) int {
    dp := make([]int, n)
    for i := 0; i < len(dp); i++ {dp[i] = 1}
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[j] = dp[j] + dp[j-1]
        }
    }
    return dp[len(dp)-1]
}


/*
    approach: top down with memoization
    - optimization of brute force dfs
    - when solving top down using memoization, we need a memo data structure.
    - this memo data structure varies from problem to problem
    - the memo data structure acts as a cache and holds values of already solved repeated problems
    - in this question, we will use a matrix as our memo cache
    
    time: o(mn)
    space: o(max(m,n)) for recursion stack + o(mn) memo matrix
*/
// func uniquePaths(m int, n int) int {
//     memo := make([][]int, m+1)
//     for i := 0; i < len(memo); i++ {memo[i] = make([]int, n+1)}
    
//     var dfs func(r, c int) int
//     dfs = func(r, c int) int {
//         // base
//         if r == m || c == n {return 0}
//         if r == m-1 && c == n-1 {return 1}
        
//         // logic
//         if memo[r+1][c] == 0 {
//             memo[r+1][c] = dfs(r+1,c)
//         }
//         if memo[r][c+1] == 0 {
//             memo[r][c+1] = dfs(r, c+1)
//         }
//         return memo[r][c+1] + memo[r+1][c]
//     }
//     return dfs(0,0)
// }


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