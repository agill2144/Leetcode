func hasPath(maze [][]int, start []int, destination []int) bool {
    m := len(maze)
    n := len(maze[0])
    
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {visited[i] = make([]bool, n)}

    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    sr, sc := start[0], start[1]
    dr, dc := destination[0], destination[1]
    visited[sr][sc] = true
    q := [][]int{start}

    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr, cc := dq[0], dq[1]
        if cr == dr && cc == dc {return true}
        for _, dir := range dirs {
            nr := cr+dir[0]
            nc := cc+dir[1]
            for nr < m && nr >= 0 && nc < n && nc >= 0 && maze[nr][nc] != 1 {
                nr += dir[0]
                nc += dir[1]
            }
            nr -= dir[0]
            nc -= dir[1]
            if visited[nr][nc] {continue}
            q = append(q, []int{nr, nc})
            visited[nr][nc] = true
        }
    }
    return false
}