func isToeplitzMatrix(matrix [][]int) bool {
    m := len(matrix)
    n := len(matrix[0])
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            prev := matrix[i][j]
            if j-1 >= 0 {prev = matrix[i-1][j-1]}
            if matrix[i][j] != prev {return false}
        }
    }
    return true
}