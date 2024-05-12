func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    out := make([][]int, n-2)
    for i := 0; i < len(out); i++ {out[i] = make([]int, n-2)}

    for i := 0; i+2 < n ; i++ {
        for j := 0; j+2 < n; j++ {
            endRow := i+2
            endCol := j+2
            maxVal := math.MinInt64
            for ii := i; ii <= endRow; ii++{
                for jj := j; jj <= endCol; jj++ {
                    maxVal = max(maxVal, grid[ii][jj])
                }
            }

            out[i][j] = maxVal

        }
    }    

    return out
}