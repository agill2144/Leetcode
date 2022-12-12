/*
    approach 1: linear search
    - iterate thru entire matrix until you find target
    time: o(mn)
    space: o(1)
*/
// func searchMatrix(matrix [][]int, target int) bool {    
//     m := len(matrix)
//     n := len(matrix[0])
    
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if matrix[i][j] == target {
//                 return true
//             }
//         }
//     }
//     return false
// }

/*
    approach 2: for each row, run a binary search if target lies within that row
    - Loop over rows
    - if target is within this row ( target >= 0th col and target <= last col)
        - then binary search this array
    
    time: o(m) * o(logn)
*/
// func searchMatrix(matrix [][]int, target int) bool {    
//     m := len(matrix)
//     n := len(matrix[0])
    
//     for i := 0; i < m; i++ {
//         if target >= matrix[i][0] && target <= matrix[i][n-1] {
//             left := 0
//             right := n-1
//             for left <= right {
//                 mid := left + (right-left)/2
//                 if matrix[i][mid] == target {
//                     return true
//                 } else if matrix[i][mid] > target {
//                     right = mid-1
//                 } else {
//                     left = mid+1
//                 }
//             }
//             return false
//         }
//     }
//     return false
// }


/*
    approach 3: Two pass; binary search the row and then binary search that row array only
    - optimization of above solution
    - first, we will use binary search to find the row
    - once row is discovered, then we will use binary search to find the element
    
    time: o(logm) + o(logn)
    space: o(1)
*/
// func searchMatrix(matrix [][]int, target int) bool {    
//     m := len(matrix)
//     n := len(matrix[0])
    
//     left := 0
//     right := m-1
//     rowIdx := -1
//     for left <= right {
//         mid := left + (right-left)/2
//         if target >= matrix[mid][0] && target <= matrix[mid][n-1] {
//             rowIdx = mid
//             break
//         } else if target > matrix[mid][n-1] {
//             left = mid+1
//         } else {
//             right = mid-1
//         }
//     }
//     if rowIdx == -1 {return false} // early exit
    
//     l := 0
//     r := n-1
//     for l <= r {
//         mid := l + (r-l)/2
//         if matrix[rowIdx][mid] == target {
//             return true
//         } else if target > matrix[rowIdx][mid] {
//             l = mid+1
//         } else {
//             r = mid-1
//         }
//     }
//     return false
// }


/*
    approach 4: treat the matrix as a flat 1d array and binary search
    time: o(logmn)
    space: o(1)
*/
func searchMatrix(matrix [][]int, target int) bool {    
    
    m := len(matrix)
    n := len(matrix[0])
    
    left := 0
    right := m*n
    
    for left <= right {
        mid := left + (right-left)/2
        r := mid/n // shouldnt this be mid%m ? we want to get relative row idx within the number of rows... 
        c := mid%n 
        if r >= m || c >= n {return false}
        if matrix[r][c] == target {
            return true
        } else if matrix[r][c] < target {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return false
    
}

/*
    approach 4: stair case search
    time: o(m+n)
    space: o(1)
*/
// func searchMatrix(matrix [][]int, target int) bool {
//     m := len(matrix)
//     n := len(matrix[0])
    
//     // pick a corner from which you can make a CONCRETE VALID decision from
//     // a decision to go in either direction
//     // we can pick top right or bottom left corner
//     // from both of these corners, we can make a valid decision
//     r := m-1
//     c := 0
//     for r >= 0 && c < n {
//         if matrix[r][c] == target {
//             return true
//         } else if target > matrix[r][c] {
//             c++
//         } else {
//             r--
//         }
//     }
//     return false
// }