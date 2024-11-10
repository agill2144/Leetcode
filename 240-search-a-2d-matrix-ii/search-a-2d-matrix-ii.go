func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        if target >= matrix[i][0] && target <= matrix[i][n-1] {
            for j := 0; j < n; j++ {
                if matrix[i][j] == target {return true}
            }
        }
    }
    return false
}
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