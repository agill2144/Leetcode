
func setZeroes(matrix [][]int)  {
    m := len(matrix)
    n := len(matrix[0])
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                for _, dir := range dirs {
                    r := i+dir[0]
                    c := j+dir[1]
                    for r >= 0 && r < m && c >= 0 && c < n && matrix[r][c] != 0 {
                        matrix[r][c] = math.MinInt64
                        r += dir[0]
                        c += dir[1]
                    }
                }
            }
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == math.MinInt64 {matrix[i][j] = 0 }
        }
    }
}

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