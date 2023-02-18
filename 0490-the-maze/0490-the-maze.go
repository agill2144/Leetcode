func hasPath(maze [][]int, start []int, destination []int) bool {
    dr := destination[0]
    dc := destination[1]
    sr := start[0]
    sc := start[1]
    if sr == dr && sc == dc {return true}
    m := len(maze)
    n := len(maze[0])
    
    dirs := [][]int{{-1,0},{0,1},{1,0},{0,-1}}
    q := [][]int{{sr,sc}}
    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr := dq[0]
        cc := dq[1]
    
        for _ , dir := range dirs {
            r := cr + dir[0]
            c := cc + dir[1]
            for r < m && r >= 0 && c < n && c >= 0 && maze[r][c] != 1 {
                r += dir[0]
                c += dir[1]
            }
            r -= dir[0]
            c -= dir[1]
            if maze[r][c] != 2 {
                maze[r][c] = 2
                if r == dr && c == dc {return true}
                q = append(q, []int{r,c})
            }
        }
    }
    return false
}