func updateMatrix(mat [][]int) [][]int {
    m := len(mat)
    n := len(mat[0])
    topDirs := [][]int{{-1,0},{0,-1}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 {
                minDist :=  math.MaxInt64-10
                for _, dir := range topDirs {
                    nr := i+dir[0]
                    nc := j+dir[1]
                    if nr >= 0 && nr < m && nc >= 0 && nc < n && mat[nr][nc] < minDist {
                        minDist = mat[nr][nc]
                    }
                }
                mat[i][j] = minDist+1
            }
        }
    }
    
    bottomDirs := [][]int{{1,0},{0,1}}
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if mat[i][j] != 0 {
                minDist := math.MaxInt64-10
                for _, dir := range bottomDirs {
                    nr := i+dir[0]
                    nc := j+dir[1]
                    if nr < m && nc < n && mat[nr][nc] < minDist {
                        minDist = mat[nr][nc]
                    }
                }
                mat[i][j] = min(mat[i][j], minDist+1)
            }
        }
    }
    return mat
}

func min(x, y int) int {
    if x < y {return x}
    return y
}