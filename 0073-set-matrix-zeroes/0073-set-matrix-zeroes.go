func setZeroes(matrix [][]int)  {
    m := len(matrix)
    n := len(matrix[0])
    tmp := make([][]int, m)
    for i := 0; i < m; i++ {
        tmp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            tmp[i][j] = matrix[i][j]
        }
    }

    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                tmp[i][j] = 0                
                for _, dir := range dirs {
                    r := i+dir[0]
                    c := j+dir[1]
                    
                    for r >= 0 && r < m && c >= 0 && c < n {
                        tmp[r][c] = 0
                        r+=dir[0]
                        c+=dir[1]
                    }
                }
            }
        }
    }
    fmt.Println(tmp)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if tmp[i][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

}