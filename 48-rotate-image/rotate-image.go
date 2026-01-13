func rotate(matrix [][]int)  {
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        for j := i+1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    for i := 0; i < m; i++ {
        left := 0; right := n-1
        for left < right {
            matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
            left++
            right--
        }
    }
}