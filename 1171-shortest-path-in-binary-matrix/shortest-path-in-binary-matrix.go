func shortestPathBinaryMatrix(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    if grid[0][0] == 1 || grid[m-1][n-1] == 1 {return -1}
    level := 1
    dirs := [][]int{
        {1,0},{-1,0},
        {0,1},{0,-1},
        {1,-1},{1,1},
        {-1,-1},{-1,1},
    }
    q := [][]int{{0,0}}
    grid[0][0] = -1
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cr := dq[0]
            cc := dq[1]
            if cr == m-1 && cc == n-1 {return level}
            for _, dir := range dirs {
                nr := cr + dir[0]
                nc := cc + dir[1]
                if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 0 {
                    q = append(q, []int{nr, nc})
                    grid[nr][nc] = -1
                }
            }
            qSize--
        }
        level++
    }
    return -1
}