/*
    approach: bottom up with space optimization
    - same as bottom up dp but with 1d dp array
    - how do we check the 3 dirs we need to for each problem
        - remember dp 1d array for each new row iteration, IS THE PREVIOUS ABOVE ROW IN DP MATRIX
        - so for each ith element, we have left ( j-1 ), we have above value ( j ), and diagUp we will mantain as a running variable (initially 0)
            - because dp 1d array is acting as our prev above row, we know for sure that current d[j] is the diagonal up value for the next jth element
            - therefore save the dp[j] in a temp var before changing the value of dp[j] and promote the tmp value to diagUp for the next jth element
            - therefore diagUp is running and kept up to date at each dp[j]
    
    time : o(mn)
    space: o(n)
    
    can we do this using 3 vars? ( up, left and diagUp ) ?
    - No, by the time we finish 1 row, we would have lost previous left ,up and diagUp values
        - draw it on the board
        - and you will see it working for the 1st row,
        - as soon as you have to go next row, you have lost the previous up, left, and diagUp values
        - therefore we need an array of previous / above row (i-1) to process an ith row
        - thereofore an array is needed
        
*/
func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    maxSize := 0
    dp := make([]int, n+1)
    diagUp := 0    
    for i := 1; i < m+1; i++ {
        for j := 1; j < n+1; j++ {
            newDiagUp := dp[j]
            if matrix[i-1][j-1] == '1' {
                dp[j] = min(dp[j-1], min(dp[j],diagUp))+1
            } else {
                dp[j] = 0
            }
            if dp[j] > maxSize {maxSize = dp[j]}
            diagUp = newDiagUp
        }
    }
    return maxSize*maxSize
}



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
// func maximalSquare(matrix [][]byte) int {
//     m := len(matrix)
//     n := len(matrix[0])
//     dp := make([][]int, m+1)
//     for idx, _ := range dp {
//         dp[idx] = make([]int, n+1)
//     }
//     max := 0
//     for i := 1; i < len(dp); i++ {
//         for j := 1; j < len(dp[0]); j++ {
//             if matrix[i-1][j-1] == '1' {
//                 minVal := min(dp[i-1][j-1], min(dp[i-1][j],dp[i][j-1]))
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

func min(x, y int) int{
    if x < y {return x}
    return y
}