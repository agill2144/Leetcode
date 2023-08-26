/*
    staircase search
    - pick a corner where we can make a absolute guranteed decision ( to go in either or direction )
    - until we find the target
    time = o(m+n)
    space = o(1)
    
    how is time m+n
    worst case we go all the way right of a row ( n cells)
    then we go all the way up to 0th row ( m rows )
    n cells first then m rows ; n+m
*/
// func searchMatrix(matrix [][]int, target int) bool {
//     m := len(matrix)
//     n := len(matrix[0])
    
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

/*
    binary search entire matrix once
    - pretend like the entire matrix is a 1D array
    - search for target
    - to get mid, we knew last idx in a 1D array
    - to get last idx in a matrix, it would be total elements -1 ; i.e (mn)-1
    - then; we can get mid idx left(0) + (right(mn-1)-left(0))/2
    - we need to translate mid idx to row and col
        row = mid / n
        col = mid % m
        
        - How did we arrive at the above
        - If the matrix size is 4*4 ( 4 rows and 4 cols )
        - that means there are 16 elements

        - lets say we are given 16 size 1d array ( flattened array )
        - and we have to create a matrix out of it
        - There are many ways to create a matrix out of 16 size 1d array
            - 1 element in each row
            - 2 element in each row
            - 3 element in each row
            - etc
        - Now lets say, we are told that the number of elements in each row MUST be 4, then
            - our only option is to put 4 elements in each row
            - therefore we will have 4 rows in total and each row contains 4 elements 
        - How many rows will exist if total size of 1d array is 16 and each row must have 4 elements ?
            - 4 rows
            - how ? 16/4
        - How many rows will exist if total size of 1d array is 30 and each row must have 5 elements ?
            - 30 elements / 5 elements in each row
            - 6 rows ( 30 / 5 )
        
        - Similaryly if we have a mid position,
            [10,11,12,13]
            [15,16,19,20]
            [23,30,34,60]
            [61,66,74,84]
        - total 1d array size would be 4*4 = 16
        - mid will be ( assumuming left = 0, right = 15 ) ; 15/2 = 7
        - pretend like this is the total number of elements we have to create a matrix out of give 4 elements in each row
        - 7 total elements and 4 in each row, how many rows ? (round it)
        - 7/4 = 1
        - 1 is now the row idx in the original matrix
        
        - To get a col idx from mid of a 1d array, we need a number relative to the size of the elements in the row
        - therefore c = mid % n
        - c = 7 % 4 = 3
        
        r = 1; c = 3
        now mid element in our matrix is 20 ( do the classic binary search )
    
    
    time: o(logmn)
    space: o(1)
*/
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    total := m*n
    left := 0
    right := total-1
    for left <= right {
        mid := left+(right-left)/2
        row := mid / n
        col := mid % n
        if row >= m || col >= n {return false}
        if matrix[row][col] == target {
            return true
        } else if target > matrix[row][col] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return false
}


// now binary search row and then binary search col
// time = o(logM) * o(logN)
// space = o(1)
// func searchMatrix(matrix [][]int, target int) bool {
//     m := len(matrix)
//     n := len(matrix[0])
//     l := 0
//     r := m-1
//     for l <= r {
//         rowMid := l+(r-l)/2
//         firstInRow := matrix[rowMid][0]
//         lastInRow := matrix[rowMid][n-1]
//         if target >= firstInRow && target <= lastInRow {
//             left := 0
//             right := n-1
//             for left <= right {
//                 mid := left +(right-left)/2
//                 if matrix[rowMid][mid] == target {return true}
//                 if target > matrix[rowMid][mid] {
//                     left = mid+1
//                 } else {
//                     right = mid-1
//                 }
//             }
//             return false
//         } else if target > lastInRow {
//             l = rowMid+1
//         } else {
//             r = rowMid-1
//         }
//     }
//     return false
// }

// for each row, check if target is within row range
// then binary search in that row
// time = o(m) * logn
// space = o(1)
// func searchMatrix(matrix [][]int, target int) bool {
//     m := len(matrix)
//     n := len(matrix[0])
//     for i := 0; i < m; i++ {
//         first := matrix[i][0]
//         last := matrix[i][n-1]
//         if target >= first && target <= last {
//             left := 0
//             right := n-1
//             for left <= right {
//                 mid := left +(right-left)/2
//                 if matrix[i][mid] == target {return true}
//                 if target > matrix[i][mid] {
//                     left = mid+1
//                 } else {
//                     right = mid-1
//                 }
//             }
//             // if binary searching this row did not work, this element does not exist
//             return false
//         }
//     }
//     return false
// }

// linear search 
// time = o(mn)
// space = o(1)
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
