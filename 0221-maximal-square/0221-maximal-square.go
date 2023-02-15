func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    maxSize := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                size := 1
                flag := true
                for i+size < m && j + size < n && flag {
                    
                    // go up in rows
                    for k := i+size; k>= i; k-- {
                        if matrix[k][j+size] != '1'{
                            flag = false
                            break
                        } 
                    }
                    
                    // go left in col
                    for k := j+size; k >= j; k-- {
                        if matrix[i+size][k] != '1' {
                            flag = false
                            break
                        }
                    }
                    
                    if flag {size++}
                }
                if size > maxSize {maxSize = size}
            }
        }
    }
    return maxSize*maxSize
}