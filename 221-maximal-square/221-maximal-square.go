/*
    approach: brute force
    - For each 1 we run into, we will try to form the largest square
    - How do we do that?
        - We will do diagonally down right first
        - If the this value is a 1, then we need to check 2 things
            1. From this cell, go up as far as we can in the same col UNTIL we hit a ceiling row - the ceiling in this case is the ith idx
            2. From this cell, go left as far as we can in tge same row UNTIL we hit a ceiling col - the ceiling col is the jth idx
        - If we were able to successfully do this, then continue diagonal down
        - Otherwise explore the next 1 
*/

// func maximalSquare(matrix [][]byte) int {
//     m := len(matrix)
//     n := len(matrix[0])
//     max := 0
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if matrix[i][j] == '1' {
//                 l := 1
//                 flag := true
//                 for i + l < m && j + l < n && flag {
                    
//                     for k := i+l; k >= i; k--{
//                         if matrix[k][j+l] == '0' {
//                             flag = false
//                             break
//                         }
//                     }
//                     for k := j+l; k >= j; k--{
//                         if matrix[i+l][k] == '0' {
//                             flag = false
//                             break
//                         }
//                     }
                    
//                     if flag { l++ }
//                 }
//                 if l*l > max { max = l*l }
//             }
//         }
//     }
//     return max
// }


/*
    approach: top down dp
    - we can create a dp matrix of m+1 and n+1 size ( with a dummy row and col )
    - and then we can start from the bottom of the dp matrix ( 2nd last row and 2nd last col )
    - and solve the original problem for each small subproblem
        - Orginial problem : find the largest square with 1's
        - Subproblem : a specific cell that has '1' 
        - Decisions : since we are doing top down, we are going to look right of the matrix
                and get min of 3 neighboring cells. bottom right, bottom, right ( thats why we took a dummy row and dummy col to handle edge case when going out of bounds )
                - get min and add +1 to this subproblem cell
        - Then save the max each time you fill a dp cell
    
    time : o(mn)
    space: o(mn) - dp matrix
*/
// func maximalSquare(matrix [][]byte) int {
//     m := len(matrix)
//     n := len(matrix[0])
//     dp := make([][]int, m+1)
//     for idx, _ := range dp {
//         dp[idx] = make([]int, n+1)
//     }
//     max := 0
//     for i := m-1; i >= 0; i-- {
//         for j := n-1; j >= 0; j-- {
//             if matrix[i][j] == '1' {
//                 minVal := min(dp[i+1][j+1], min(dp[i+1][j],dp[i][j+1]))
//                 dp[i][j] = minVal + 1
//                 if dp[i][j] > max {
//                     max = dp[i][j]
//                 } 
//             }
//         }
//     } 
//     return max*max
// }



/*
    approach: bottom up
    - we can create a dp matrix of m+1 and n+1 size ( with a dummy row and col )
    - and then we can start from the start of the dp matrix ( 1st row and 1st col )
    - and solve the original problem for each small subproblem
        - Orginial problem : find the largest square with 1's
        - Subproblem : a specific cell that has '1' 
        - Decisions : since we are doing bottom up, we are going to look back/left of the matrix
                and get min of 3 neighboring cells. up left, up, left ( thats why we took a dummy row and dummy col to handle edge case when going out of bounds )
                - get min and add +1 to this subproblem cell
        - Then save the max each time you fill a dp cell
    
    time : o(mn)
    space: o(mn) - dp matrix
*/
func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int, n+1)
    }
    max := 0
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if matrix[i-1][j-1] == '1' {
                minVal := min(dp[i-1][j-1], min(dp[i-1][j],dp[i][j-1]))
                dp[i][j] = minVal + 1
                if dp[i][j] > max {
                    max = dp[i][j]
                } 
            }
        }
    } 
    return max*max
}

func min(x, y int) int{
    if x < y {return x}
    return y
}