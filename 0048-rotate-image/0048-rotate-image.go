func rotate(matrix [][]int)  {
    n := len(matrix)
    cp := make([][]int, n)
    for i := 0; i < n; i++ {cp[i] = make([]int, n)}

    for i := 0;  i < n; i++ {
        for j := 0; j < n; j++ {
            cp[j][n-i-1] = matrix[i][j]
        }
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            matrix[i][j] = cp[i][j]
        }
    }
}