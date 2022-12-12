func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    
    // for i := 1; i < m; i++ { grid[i][0] += grid[i-1][0]}
    // for j := 1; j < n; j++ { grid[0][j] += grid[0][j-1]}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 {
                jVal := 0
                if j-1 >= 0 {jVal = grid[i][j-1]}
                grid[i][j] += jVal
                continue
            }
            
            if j == 0 {
                iVal := 0
                if i-1 >= 0 {
                    iVal = grid[i-1][j]
                }
                grid[i][j] += iVal
                continue
            }
            
            grid[i][j] += min(grid[i-1][j], grid[i][j-1])  
        }
    }
    return grid[m-1][n-1]
}
func min(x, y int) int {
    if x < y {return x}
    return y
}