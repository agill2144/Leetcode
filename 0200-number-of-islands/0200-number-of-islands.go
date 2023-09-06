/*
    this problem is just finding number of connected components
    - we can do this in place
    - or use a visited data structure ( like bool matrix ) ; just like all other graph problems
    
    
    can be done using dfs
    can be done using bfs
    - at each cell with value == 1
        - run dfs/bfs
        - increment your counter
        - dfs/bfs will go mark all the cells in that component visited

    time = o(mn)
    space = o(mn) in dfs we could end up adding all cells in our recursive stack
*/
func numIslands(grid [][]byte) int {
    // in place dfs 
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    m := len(grid)
    n := len(grid[0])
    
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] != '1' {return}
        
        // logic
        grid[r][c] = '0'
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                dfs(i, j)
                count++
            }
        }
    }
    return count
}