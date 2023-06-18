/*
    identical to : https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
    instead of finding the longest one, count all the paths that exist
*/
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
                memo[r][c] += dfs(nr, nc)
                memo[r][c] %= mod
            }
        }
        return memo[r][c]
    }
    

    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            total += dfs(i, j)
            total %= mod
        }
    }
    return total
}