func isToeplitzMatrix(matrix [][]int) bool {
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            curr := matrix[i][j]
            prev := curr
            if i-1 >= 0 && j-1 >= 0 {prev = matrix[i-1][j-1]}
            if curr != prev {return false}
        }
    }
    return true
}