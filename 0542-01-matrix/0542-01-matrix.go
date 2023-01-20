func updateMatrix(mat [][]int) [][]int {
    m := len(mat)
    n := len(mat[0])
    topLeftDirs := [][]int{{0,-1},{-1,0}}
    bottomRightDirs := [][]int{{1,0},{0,1}}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {continue}
            minVal := math.MaxInt64-10
            for _, dir := range topLeftDirs {
                r := i+dir[0]
                c := j+dir[1]
                if r >= 0 && c >= 0 && mat[r][c] < minVal {
                    minVal = mat[r][c]
                }
            }
            mat[i][j] = minVal+1
        }
    }
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if mat[i][j] == 0 {continue}
            minVal := math.MaxInt64-10
            for _, dir := range bottomRightDirs {
                r := i+dir[0]
                c := j+dir[1]
                if r < m && c < n && mat[r][c] < minVal {
                    minVal = mat[r][c]
                }
            }
            mat[i][j] = min(mat[i][j], minVal+1)
        }
    }
    return mat
    
}

func min(x, y int) int {
    if x < y {return x}
    return y
}