// time: o(mn)
func hasPath(maze [][]int, start []int, destination []int) bool {
    dr := destination[0]
    dc := destination[1]
    sr := start[0]
    sc := start[1]
    if sr == dr && sc == dc {return true}
    m := len(maze)
    n := len(maze[0])
    
    dirs := [][]int{{-1,0},{0,1},{1,0},{0,-1}}
    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        // base
        if maze[r][c] == 2 {return false}
        if r == dr && c == dc {return true}
        
        // logic
        maze[r][c] = 2
        for _ , dir := range dirs {
            nr := r + dir[0]
            nc := c + dir[1]
            for nr < m && nr >= 0 && nc < n && nc >= 0 && maze[nr][nc] != 1 {
                nr += dir[0]
                nc += dir[1]
            }
            nr -= dir[0]
            nc -= dir[1]
            if ok := dfs(nr, nc); ok {return true}
        }
        
        return false
    }
    return dfs(sr, sc)
}