func numDistinctIslands(grid [][]int) int {
    dirs := [][]interface{}{{1,0,"D"},{-1,0,"U"},{0,1,"R"},{0,-1,"L"}}
    m := len(grid)
    n := len(grid[0])
    set := map[string]int{}
    
    var dfs func(r, c int, path *string)
    dfs = func(r, c int, path *string) {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] != 1 {return}
        
        // logic
        grid[r][c] = -1
        for _, dir := range dirs {
            *path = *path+dir[2].(string)
            dfs(r+dir[0].(int), c+dir[1].(int), path)
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                path := ""
                dfs(i,j,&path)
                set[path]++
            }
        }
    }
    return len(set)
}