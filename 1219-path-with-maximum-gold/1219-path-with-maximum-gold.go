func getMaximumGold(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    max := 0
    var backtrack func(r, c int, total int) 
    dirs := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    backtrack = func(r, c int, total int) {
        // base
        if r == m || r < 0 || c == n || c < 0 || grid[r][c] <= 0 {
            return
        }
        
        // logic
        total += grid[r][c]
        // action
        grid[r][c] *= -1
        if total > max {max = total}
        for _, dir := range dirs {
            // recurse
            backtrack(r+dir[0], c+dir[1], total)
        }
        // backtrack
        grid[r][c] *= -1
        
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != 0 {
                backtrack(i,j,0)
            }
        }
    }
    return max
}