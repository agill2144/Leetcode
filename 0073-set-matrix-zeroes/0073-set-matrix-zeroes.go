func setZeroes(matrix [][]int)  {
    m := len(matrix)
    n := len(matrix[0])
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                for _, dir := range dirs {
                    r := i+dir[0]
                    c := j+dir[1]
                    for r >= 0 && r < m && c >= 0 && c < n {
                        if matrix[r][c] != 0 {
                            matrix[r][c] = math.MaxInt64
                        }
                        r += dir[0]
                        c += dir[1]
                    }
                }
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == math.MaxInt64 {matrix[i][j] = 0}
       }
    }
}