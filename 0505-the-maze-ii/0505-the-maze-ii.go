func shortestDistance(maze [][]int, start []int, destination []int) int {
    m := len(maze)
    n := len(maze[0])
    
    // why dist matrix? and not use the same maze to track whats visited or not
    // this is like a graph problem, where we can reach a node/cell with 2 different distances
    // we may reach a cell with some distance, but some other path leads to a smaller distance
    // i.e we have to re-process a node again if we reached at a smaller distance ( until queue is empty )
    dist := make([][]int, m)
    for i := 0; i < m; i++ {
        dist[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dist[i][j] = math.MaxInt64
        }
    }
    sr, sc := start[0], start[1]
    dr, dc := destination[0], destination[1]

    dirs := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    q := [][]int{{sr,sc,0}}
    dist[sr][sc] = 0
    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr := dq[0]
        cc := dq[1]
        currentDist := dq[2]
        
        for _, dir := range dirs {
            nr := cr + dir[0]
            nc := cc + dir[1]
            neiDist := 0
            for nr >= 0 && nr < m && nc >= 0 && nc < n && maze[nr][nc] != 1 {
                neiDist++
                nr += dir[0]
                nc += dir[1]
            }
            nr -= dir[0]
            nc -= dir[1]
            totalDist := currentDist + neiDist
            if totalDist < dist[nr][nc] {
                dist[nr][nc] = totalDist
                q = append(q, []int{nr, nc, totalDist})
            }
        }
    }
    if dist[dr][dc] == math.MaxInt64 {return -1}
    return dist[dr][dc]
}