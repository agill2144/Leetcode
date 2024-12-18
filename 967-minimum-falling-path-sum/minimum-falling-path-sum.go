func minFallingPathSum(matrix [][]int) int {
    n := len(matrix)
    dirs := [][]int{{-1,0},{-1,-1},{-1,1}}
    for i := 1; i < n; i++ {
        for j := 0; j < n; j++ {
            minVal := math.MaxInt64-1000
            for _, dir := range dirs {
                nr := i+dir[0]
                nc := j+dir[1]
                if nr >= 0 && nc < n && nc >= 0 {
                    minVal = min(minVal, matrix[nr][nc])
                }
            }
            matrix[i][j] += minVal
        }
    }
    out := math.MaxInt64
    for j := 0; j < n; j++ {
        out = min(out, matrix[n-1][j])
    }
    return out
}