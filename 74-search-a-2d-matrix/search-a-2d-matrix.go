/*
    approach: binary search within each row
    - loop over all the rows
    - if our target lies in a ith row 
        - then binary search in this row ( n elements )
        - therefore; logn
    
    total time = o(m) + logn
    space = o(1)
*/
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        if target >= matrix[i][0] && target <= matrix[i][n-1] {
            left := 0
            right := n-1
            // vanilla binary search on a 1D array
            for left <= right {
                mid := left + (right-left)/2
                if matrix[i][mid] == target {return true}
                if target > matrix[i][mid] {
                    left = mid+1
                } else {
                    right = mid-1
                }
            }
        }
    }
    return false
}
/*
    approach: brute force
    - linear scan the entire matrix

    time = o(mn)
    space = o(1)
*/
// func searchMatrix(matrix [][]int, target int) bool {
//     m := len(matrix)
//     n := len(matrix[0])
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if matrix[i][j] == target {return true}
//         }
//     }
//     return false
// }