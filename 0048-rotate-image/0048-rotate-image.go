func rotate(matrix [][]int)  {
    m := len(matrix)
    n := len(matrix[0])
    tmp := make([][]int, m)
    for i := 0; i < m; i++ {
        tmp[i] = make([]int, n)
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            tmp[j][n-i-1] = matrix[i][j]
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            matrix[i][j] = tmp[i][j]
        }
    }
    
}