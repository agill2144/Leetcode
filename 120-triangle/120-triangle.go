func minimumTotal(triangle [][]int) int {
    m := len(triangle)
    out := math.MaxInt64
    
    dirs := [][]int{{-1,0},{-1,-1}}
    for i := 1; i < m; i++ {
        for j := 0; j < len(triangle[i]); j++ {
            minVal := math.MaxInt64
            for _, dir := range dirs {
                r := i+dir[0]
                c := j+dir[1]
                if r >= 0 && c >= 0 && c < len(triangle[i-1]) {
                    minVal = min(minVal, triangle[r][c])
                }
            }
            triangle[i][j] += minVal
        }
    }
    
    for j := 0; j < len(triangle[m-1]); j++  {
        out = min(out, triangle[m-1][j])
    }
    return out
}

func min(x, y int) int {
    if x < y { return x }
    return y
}