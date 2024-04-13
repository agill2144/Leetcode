func minimumTotal(triangle [][]int) int {
    dirs := [][]int{{-1,0},{-1,-1}}
    m := len(triangle)
    out := math.MaxInt64
    for i := 1; i < m; i++ {
        prevRowLen := len(triangle[i-1])
        for j := 0; j < len(triangle[i]); j++ {
            // look up and get min val
            minVal := math.MaxInt64
            for _, dir := range dirs {
                nr := i+dir[0]
                nc := j+dir[1]
                if nr >= 0 && nc >= 0 && nc < prevRowLen {
                    minVal = min(minVal, triangle[nr][nc])
                }
            }
            triangle[i][j] += minVal

            if i == m-1 {
                out = min(out, triangle[i][j])
            }
        }
    }
    if out == math.MaxInt64 {
        for j := 0; j < len(triangle[m-1]); j++  {
            out = min(out, triangle[m-1][j])
        }
    return out        
    }
    return out
}