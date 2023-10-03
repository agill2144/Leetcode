func rotate(matrix [][]int)  {
    n := len(matrix)
    tmp := make([][]int, n)
    for i := 0; i < n; i++ {
        tmp[i] = make([]int, n)
    }
    c := n-1
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            tmp[j][c] = matrix[i][j]
        }
        c--
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            matrix[i][j] = tmp[i][j]
        }
    } 
}