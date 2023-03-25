func numDistinctIslands(grid [][]int) int {
    // the way to store distinct islands is by storing the path we took
    paths := map[string]struct{}{}
    dirs := [][]interface{}{{0,1,"R"},{1,0,"D"},{0,-1,"L"},{-1,0,"U"}}
    
    m := len(grid)
    n := len(grid[0])
    
    var dfs func(r, c int, path *string)
    dfs = func(r, c int, path *string){
        // base
        
        
        // logic
        grid[r][c] = 0
        for _, dir := range dirs {
            nr := dir[0].(int) + r
            nc := dir[1].(int) + c
            if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
                *path += dir[2].(string)
                dfs(nr, nc, path)
            }
            *path += "E"
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                path := "S"
                dfs(i,j, &path)
                paths[path] = struct{}{}
            }
        }
    }
    return len(paths)
}