func setZeroes(matrix [][]int)  {
    m := len(matrix)
    n := len(matrix[0])
    // collect rows and cols that should be replaced with 0's
    rows := make([]bool, m)
    cols := make([]bool, n)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                rows[i] = true
                cols[j] = true
            }
        }
    }
    
    // replace cell val is cell is in either row or col that should be replaced with 0's
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rows[i] || cols[j] {
                matrix[i][j] = 0
            }
        }
    }
}