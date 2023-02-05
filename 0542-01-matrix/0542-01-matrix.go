func updateMatrix(mat [][]int) [][]int {
    m := len(mat)
    n := len(mat[0])
    topLeftDirs := [][]int{{-1,0},{0,-1}}
    bottomRight := [][]int{{1,0},{0,1}}

    for  i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {continue}

            minVal := math.MaxInt64-100
            for _, dir := range topLeftDirs {
                nr := i+dir[0]
                nc := j+dir[1]
                if nr >= 0 && nc >= 0 && mat[nr][nc] < minVal {
                    minVal = mat[nr][nc]
                }
            }
            mat[i][j] = minVal + 1
        }
    }
    
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if mat[i][j] == 0 {continue}
            
            minVal := math.MaxInt64-100
            for _ , dir := range bottomRight {
                nr := i+dir[0]
                nc := j+dir[1]
                if nr < m && nc < n && mat[nr][nc] < minVal {
                    minVal = mat[nr][nc]
                }
            }
            minVal += 1
            
            mat[i][j] = min(mat[i][j], minVal)
        }
    }
    

    return mat
}

func min(x, y int) int {
    if x < y {return x}
    return y
}