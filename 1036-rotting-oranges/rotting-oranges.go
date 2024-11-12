func orangesRotting(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    fresh := 0
    q := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {fresh++}
            if grid[i][j] == 2 {q = append(q, []int{i, j})}
        }
    }
    if fresh == 0 {return 0}
    if len(q) == 0 {return -1}
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    time := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cr := dq[0]
            cc := dq[1]
            for _, dir := range dirs {
                nr := cr + dir[0]
                nc := cc + dir[1]
                if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
                    fresh--
                    q = append(q, []int{nr, nc})
                    grid[nr][nc] = 2
                } 
            }
            qSize--
        }
        time++
    }
    if fresh > 0 {return -1}
    return time-1
}