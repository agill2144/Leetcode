func countPaths(grid [][]int) int {
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    m := len(grid)
    n := len(grid[0])
    total := 0
    mod := 1000000007
    // all vals will be 0; i.e not-yet-solved
    memo := make([][]int, m)
    for i := 0; i < m; i++ {memo[i] = make([]int, n)}
    

    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        // base
        // already solved? return the solved value
        if memo[r][c] != 0 {
            return memo[r][c]
        }
        
        // logic
        // otherwise solve for this cell
        // initially the count will be 1
        memo[r][c] = 1
        for _, dir := range dirs {
            nr := r+dir[0]
            nc := c+dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] > grid[r][c] {
                // if the neighboring cell is inbound
                // and its value > current cell
                // recurse and find out best ans for neighbor
                // the +1 is part of the recursive call because
                // the neighbor cell is increasing compared to current
                // so if current cell longest-increasing path is of size 1 and
                // we discovered a new adjacent cell to explore, the new size will be currentCount + 1
                memo[r][c] = (memo[r][c] + dfs(nr, nc)) % mod
            }
        }
        return memo[r][c]
    }
    

    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            total = (total + dfs(i, j) ) % mod
        }
    }
    return total
}