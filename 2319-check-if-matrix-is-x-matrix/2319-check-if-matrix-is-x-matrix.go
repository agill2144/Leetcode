func checkXMatrix(grid [][]int) bool {
    n := len(grid)
    r := 0
    c := 0
    for r < n && c < n {
        if grid[r][c] == 0 {return false}
        grid[r][c] *= -1 // visited
        r++; c++
    }
    
    r = 0; c = n-1
    for r < n && c >= 0 {
        if grid[r][c] == 0 {return false}
        if grid[r][c] > 0 {
            grid[r][c] *= -1
        }
        r++; c--
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < n ; j++ {
            if grid[i][j] > 0 {return false}
        }
    }
    return true
}