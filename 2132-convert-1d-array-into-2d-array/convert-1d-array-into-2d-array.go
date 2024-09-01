func construct2DArray(original []int, m int, n int) [][]int {
    if len(original) != m*n {return nil}
    ptr := 0
    matrix := make([][]int, m)
    for i := 0; i < len(matrix); i++ {
        matrix[i] = make([]int,n)
        for j := 0; j < len(matrix[i]); j++ {
            matrix[i][j] = original[ptr]
            ptr++
        }
    }
    return matrix
}