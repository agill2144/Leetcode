func rotate(matrix [][]int)  {
    n := len(matrix)
    // transpose
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n/2; j++ {
            matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
        }        
    }

}

// func rotate(matrix [][]int)  {
//     n := len(matrix)
//     tmp := make([][]int, n)
//     for i := 0; i < n; i++ {
//         tmp[i] = make([]int, n)
//     }
    
//     for i := 0; i < n; i++ {
//         for j := 0; j < n; j++ {
//             tmp[j][n-1-i] = matrix[i][j]
//         }
//     }
    
//     for i := 0; i < n; i++ {
//         for j := 0; j < n; j++ {
//             matrix[i][j] = tmp[i][j]
//         }
//     } 
// }