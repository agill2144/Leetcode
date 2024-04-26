func minFallingPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    ans := math.MaxInt64
    
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {        
            minSum := math.MaxInt64
            for k := 0; k < n; k++ {
                if k == j {continue}
                minSum = min(minSum, grid[i-1][k] + grid[i][j])
            }
            grid[i][j] = minSum
        }
    }
    for j := 0; j < n; j++ {
        ans = min(ans, grid[m-1][j])
    }
    return ans
}

func min(x, y int) int {
    if x < y {return x}
    return y
}