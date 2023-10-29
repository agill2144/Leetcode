func rotate(matrix [][]int)  {
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        for j := i+1; j < n; j++ {
            matrix[i][j],matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n/2; j++ {
            matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
        }
    }
    
}