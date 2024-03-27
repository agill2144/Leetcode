/*
    approach: optimal binary search
    - pretend like the entire matrix is a 1D array
    - and binary search on this imaginary 1D array index
    - total elements = m * n
    - therefore left = 0 and right = total-1
    - we can find mid idx.
    - how do we translate this mid idx of our imaginary array to row/col positions
    - for row;
        - if mid idx = 5
        - this means there are 5 elements before this 5th idx; at indicies 0,1,2,3,4
        - so there are supposed to 5 elements on the left side of this 5th index
        - how many elements do we have in a row ?
            - n
            - lets say n = 4
            - this means in the 0th row there are 4 total elements
            - which also means 5th element will be in row 1
        - another way to think about this is to ask
            - if have 5 elements to place in a matrix ( which does not exist yet )
            - and each row can only have n elements
            - which row would have 5th element
        - therefore row = mid/n
    - for col;
        - if mid idx = 5
        - and n = 4 ( num of elements in each row or col size )
        - if 4 were in the 0th row, how many are left ? 1
        - this is remainder; mid % n
        - remainder gives us how many more elements that could not fit in above row 
        - and are now part of this row, the row that contains 5th element
        - except the first mid%n cells in this row are taken by elements before 5th element
        - therefore col = mid % n        
    - now we have translated mid idx of 1d array to row / col
    - this becomes a vanilla binary search

    time = o(log mn)
    space = o(1)
*/
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    total := m*n
    left := 0
    right := total-1
    for left <= right {
        mid := left + (right-left)/2
        row := mid / n
        col := mid % n
        if matrix[row][col] == target {return true}
        if target > matrix[row][col] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return false
}

/*
    approach: binary search within each row
    - loop over all the rows
    - if our target lies in a ith row 
        - then binary search in this row ( n elements )
        - therefore; logn
    
    total time = o(m) + logn
    space = o(1)
*/
// func searchMatrix(matrix [][]int, target int) bool {
//     m := len(matrix)
//     n := len(matrix[0])
//     for i := 0; i < m; i++ {
//         if target >= matrix[i][0] && target <= matrix[i][n-1] {
//             left := 0
//             right := n-1
//             // vanilla binary search on a 1D array
//             for left <= right {
//                 mid := left + (right-left)/2
//                 if matrix[i][mid] == target {return true}
//                 if target > matrix[i][mid] {
//                     left = mid+1
//                 } else {
//                     right = mid-1
//                 }
//             }
//         }
//     }
//     return false
// }
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