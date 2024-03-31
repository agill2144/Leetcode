func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            peak := true
            for _, dir := range dirs {
                nr, nc := i+dir[0], j+dir[1]
                if nr >= 0 && nr < m && nc >= 0 && nc < n && mat[nr][nc] > mat[i][j] {
                    peak = false
                    break
                }
            }
            if peak {
                return []int{i, j}
            }
        }
    }
    return nil
}