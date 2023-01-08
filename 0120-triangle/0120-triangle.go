func minimumTotal(triangle [][]int) int {
    m := len(triangle)
    if m == 1 {return triangle[0][0]}
    dirs := [][]int{{-1,0},{-1,-1}}
    out := math.MaxInt64
    for i := 1; i < m; i++ {
        n := len(triangle[i])
        prevN := len(triangle[i-1])
        
        for j := 0; j < n; j++ {

            minVal := math.MaxInt64
            for _, dir := range dirs{
                r := i+dir[0]
                c := j+dir[1]
                if r >= 0 && c >= 0 && c < prevN {
                    minVal = min(triangle[r][c], minVal)
                }
            }
            triangle[i][j] += minVal
            
            if i == m-1 {
                out = min(out, triangle[i][j])
            }
        }
    }
    return out
}

func min(x, y int) int {
    if x < y {return x}
    return y
}