/*
    approach: top down - matrix traversal
    - start from the the 1st row
    - and for each col;
    - look up and select min elements from 3 choices ( diag-left-up, up, diag-right-up )
    - once we have made a decision, add that decision value to the current col
    - and work all the way till the last row
    - finally return the min val from the last row
    
    time: o(mn)
    space: o(1)
*/

// func minFallingPathSum(matrix [][]int) int {
//     m := len(matrix)
//     n := len(matrix[0])
//     min := math.MaxInt64
    
//     dirs := [][]int{{-1,0},{-1,-1},{-1,1}}
//     for i := 1; i < m; i++ {
//         for j := 0; j < n; j++ {        
//             smallestChoice := math.MaxInt64
//             for _, dir := range dirs {
//                 r := i+dir[0]
//                 c := j+dir[1]
//                 if r >= 0 && c >= 0 && c < n {
//                     if matrix[r][c] < smallestChoice {
//                         smallestChoice = matrix[r][c]
//                     }
//                 }
//             }
//             matrix[i][j] += smallestChoice
//         }
//     }
//     for j := 0; j < n; j++ {
//         if matrix[m-1][j] < min {
//             min = matrix[m-1][j]
//         }
//     }
//     return min
// }


/*
    approach: bottom up - matrix traversal
    - start from the the 2nd last row
    - and for each col;
    - look up and select min elements from 3 choices ( diag-left-down, down, diag-right-down )
    - once we have made a decision, add that decision value to the current col
    - and work all the way up to the top row
    - finally return the min val from the top row
    
    time: o(mn)
    space: o(1)
*/

// func minFallingPathSum(matrix [][]int) int {
//     m := len(matrix)
//     n := len(matrix[0])
//     min := math.MaxInt64
    
//     dirs := [][]int{{1,0},{1,-1},{1,1}}
//     for i := m-2; i >= 0; i-- {
//         for j := 0; j < n; j++ {
//             smallest := math.MaxInt64
//             for _, dir := range dirs {
//                 r := dir[0]+i
//                 c := dir[1]+j
//                 if r >= 0 && r < m && c < n && c >= 0 {
//                     if matrix[r][c] < smallest {
//                         smallest = matrix[r][c]
//                     }
//                 }
//             }
//             matrix[i][j] += smallest
        
//         }
//     }
//     for j := 0; j < n; j++ {
//         if matrix[0][j] < min {
//             min = matrix[0][j]
//         }
//     }
//     return min
// }



/*
    approach: bottom up - matrix traversal - if we cannot mutate the matrix - use dp matrix
    - start from the the 2nd last row
    - and for each col;
    - look up and select min elements from 3 choices ( diag-left-down, down, diag-right-down )
    - once we have made a decision, add that decision value to the current col
    - and work all the way up to the top row
    - finally return the min val from the top row
    
    time: o(mn)
    space: o(mn)
*/

func minFallingPathSum(matrix [][]int) int {
    m := len(matrix)
    n := len(matrix[0])
    min := math.MaxInt64
    dp := make([][]int, m)
    for idx, _ := range dp {
        dp[idx] = make([]int, n)
    }
    for j := 0; j < n; j++ {
        dp[m-1][j] = matrix[m-1][j]
    }
    
    dirs := [][]int{{1,0},{1,-1},{1,1}}
    for i := m-2; i >= 0; i-- {
        for j := 0; j < n; j++ {
            smallest := math.MaxInt64
            for _, dir := range dirs {
                r := dir[0]+i
                c := dir[1]+j
                if r >= 0 && r < m && c < n && c >= 0 {
                    if dp[r][c] < smallest {
                        smallest = dp[r][c]
                    }
                }
            }
            dp[i][j] += smallest + matrix[i][j]
        
        }
    }
    for j := 0; j < n; j++ {
        if dp[0][j] < min {
            min = dp[0][j]
        }
    }
    return min
}