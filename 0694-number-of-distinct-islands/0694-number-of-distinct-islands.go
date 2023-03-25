func numDistinctIslands(grid [][]int) int {
    // the way to store distinct islands is by storing the path we took
    paths := map[string]struct{}{}
    dirs := [][]interface{}{{0,1,"R"},{1,0,"D"},{0,-1,"L"},{-1,0,"U"}}
    
    m := len(grid)
    n := len(grid[0])
    
    var dfs func(r, c int, path *strings.Builder)
    dfs = func(r, c int, path *strings.Builder){
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] != 1 {
            path.WriteString("-")
            return
        }
        
        
        // logic
        grid[r][c] = 0
        for _, dir := range dirs {
            path.WriteString(dir[2].(string))
            dfs(dir[0].(int) + r, dir[1].(int) + c, path)
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                path := new(strings.Builder)
                path.WriteString("S")
                dfs(i,j, path)
                paths[path.String()] = struct{}{}
            }
        }
    }
    return len(paths)
}