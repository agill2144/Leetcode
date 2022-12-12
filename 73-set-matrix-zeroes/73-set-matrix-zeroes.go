// func setZeroes(matrix [][]int)  { 
//     m := len(matrix)
//     n := len(matrix[0])
    
//     cp := make([][]int, m)
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if cp[i] == nil {
//                 cp[i] = make([]int, n)
//             } 
//             cp[i][j] = math.MaxInt64
//         }
//     }
    
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if matrix[i][j] == 0 {
                
//                 for c := 0; c < n; c++ {
//                     cp[i][c] = 0
//                 }
                
//                 for c := 0; c < m; c++ {
//                     cp[c][j] = 0
//                 }
//             }
//         }
//     }
    
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if cp[i][j] == 0 {
//                 matrix[i][j] = 0
//             }
//         }
//     }
    
// }

func setZeroes(matrix [][]int)  { 
    m := len(matrix)
    n := len(matrix[0])
    
    isCol := false
    for i := 0; i < m; i++ {
        if matrix[i][0] == 0 {isCol = true}
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
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
    
    if matrix[0][0] == 0 {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }
    
    if isCol {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
}