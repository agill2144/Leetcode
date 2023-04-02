func numDistinctIslands(grid [][]int) int {
    dirs := [][]interface{}{{0,-1,"L"},{1,0,"D"},{0,1,"R"},{-1,0,"U"}}
    m := len(grid)
    n := len(grid[0])
    set := map[string]struct{}{}
    
    var dfs func(r, c int, path *strings.Builder)
    dfs = func(r, c int, path *strings.Builder) {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] != 1 {return}
        
        // logic
        grid[r][c] = -1
        for _, dir := range dirs {
            path.WriteString(dir[2].(string))
            dfs(r+dir[0].(int), c+dir[1].(int), path)
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                path := new(strings.Builder)
                dfs(i,j,path)
                set[path.String()]=struct{}{}
            }
        }
    }
    return len(set)
}