func transpose(matrix [][]int) [][]int {
    m := len(matrix)
    n := len(matrix[0])
    out := make([][]int, n)
    for i := 0; i < n; i++ {out[i] = make([]int, m)}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            out[j][i] = matrix[i][j]
        }
    }
    
    return out
}