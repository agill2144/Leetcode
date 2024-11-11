func generate(numRows int) [][]int {
    matrix := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        n := i+1
        matrix[i] = make([]int, n)
        matrix[i][0] = 1
        matrix[i][n-1] = 1
        for j := 1; j < n-1; j++ {
            matrix[i][j] = matrix[i-1][j] + matrix[i-1][j-1]
        }
    }
    return matrix
}