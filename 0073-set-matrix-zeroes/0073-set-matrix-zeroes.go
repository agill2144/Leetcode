/*
    whenever we have to optimize on space, consider the 2 options
    1. can we use the output space ? 
        - we dont return anything, so no
    2. can we use the input as our space ?
        - maybe
        - instead of saving those edges in a separate hashset,
        - we will use existing left edge and top edge of the matrix to keep track of 
            which rows and cols to be replaced with 0's
    
    approach : keep track of rows and cols to be replaced in the same input matrix
    
*/
func setZeroes(matrix [][]int)  {
    colZero := false
    rowZero := false
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                if i == 0 && j == 0  {
                    rowZero = true
                    colZero = true                    
                } else if j == 0 {
                    colZero = true
                } else if i == 0 {
                    rowZero = true
                } else {
                    matrix[i][0] = 0
                    matrix[0][j] = 0
                }
            }
        }
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    if rowZero {
        for j := 0; j < n; j++ {matrix[0][j] = 0}
    }
    if colZero {
        for i := 0; i< m; i++ {matrix[i][0] = 0}
    }
    
}

/*
    approach: keep track of rows and cols to be replaced 
    - instead of marking the entire row and col with newZero value, each time we run into a 0
        - this caused o(mn * m+n)  time-complexity
        - there extra m+n for each cell
    - we can instead keep track of which row and col should be replaced with 0s
    - so first we will go thru entire matrix and identify which rows and cols should be replaced with 0s
    - and then another pass of matrix traversal, in which for each cell we will check
        - does this row i needs to replaced with 0s ? or does this col j needs to replaced with 0s ?
        - whatever ds we use to keep track of rows and cols, it should make searching easy
        - so a hashset to make searching easy
    - thinking of saving the edges of the matrix ( row edge, col edge ) in a hashset

    time = o(mn) + o(mn)
    space = o(m) + o(n)
    
    better time complexity, but still using space
*/
// func setZeroes(matrix [][]int)  {
//     m := len(matrix)
//     n := len(matrix[0])
//     rows := map[int]struct{}{}
//     cols := map[int]struct{}{}
    
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if matrix[i][j] == 0 {
//                 rows[i] = struct{}{}
//                 cols[j] = struct{}{}
//             }
//         }
//     }

//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             _, r := rows[i]
//             _, c := cols[j]
//             if r || c {
//                 matrix[i][j] = 0
//             }
//         }
//     }    
// }

/*
    The problem with blindly marking elements 0 is that
    when we run into a 0 from input matrix
    we will go thru that entire row and col and flip all values 0
    and then our iteration goes to next cell of input matrix and finds EXTRA 0's that we just put in.
    This is the side effect of replacing a value with same value while continuing to look for the same value
    
    Instead, we will come up with a new identifier that helps us avoid touching cells that are newely converted 0s
    
    approach:  in-place, 2 pass
    - 1st pass
        - when we run into a 0
        - replace the entire row and col with new ID ( a new zero ID ) = math.MinInt
        - the -Inf value acts as our new 0
        - However another case is this
        [1,2,0,3,4,0]
        [5,6,7,8,9,1]
        - when we run into our first 0 (r = 0; c = 2)
        - by processing that row and col, our matrix will look like this
        [-inf,-inf,0,-inf,-inf,-inf]
        [5,     6, 7,  8,  9,     1]
        - the new problem is we replaced a NOT-PROCESSED 0 (r=0,c=5) because of an earlier 0(r = 0,c=2)
        - therefore we should not replace existing 0's with -inf values just yet
        - Therefore skip cells that are 0's when processing a row and a col]    
    - 2nd pass
        - flip the -inf to 0

    
    time = o(mn * m+n for marking)+o(mn)
    space = o(1)
*/
// func setZeroes(matrix [][]int)  {
//     m := len(matrix)
//     n := len(matrix[0])
//     dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if matrix[i][j] == 0 {
//                 for _, dir := range dirs {
//                     r := i+dir[0]
//                     c := j+dir[1]
//                     for r >= 0 && r < m && c >= 0 && c < n && matrix[r][c] != 0 {
//                         matrix[r][c] = math.MinInt64
//                         r += dir[0]
//                         c += dir[1]
//                     }
//                 }
//             }
//         }
//     }
    
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if matrix[i][j] == math.MinInt64 {matrix[i][j] = 0 }
//         }
//     }
// }

/*
    approach: brute force
    - allocate a tmp matrix
    - deep copy input into tmp 
    - loop over input matrix
    - if we run into a 0
    - replace all the 4 dirs in tmp matrix with 0's (including this position )
    - then write back the tmp matrix to input matrix ( take all the 0's and write them back to input matrix )
    
    time = o(mn) + o(mn * m+n) + o(mn)
    space = o(mn)
*/
// func setZeroes(matrix [][]int)  {
//     m := len(matrix)
//     n := len(matrix[0])
//     tmp := make([][]int, m)
    
//     // time = o(mn)
//     for i := 0; i < m; i++ {
//         tmp[i] = make([]int, n)
//         for j := 0; j < n; j++ {
//             tmp[i][j] = matrix[i][j]
//         }
//     }

//     dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}

//     for i := 0; i < m; i++ { // m
//         for j := 0; j < n; j++ {
//             if matrix[i][j] == 0 { // n
//                 tmp[i][j] = 0                
//                 for _, dir := range dirs { //  m+n
//                     r := i+dir[0]
//                     c := j+dir[1]
                    
//                     for r >= 0 && r < m && c >= 0 && c < n {
//                         tmp[r][c] = 0
//                         r+=dir[0]
//                         c+=dir[1]
//                     }
//                 }
//             }
//         }
//     }
//     // mn time
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if tmp[i][j] == 0 {
//                 matrix[i][j] = 0
//             }
//         }
//     }
// }