func orangesRotting(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}

    fresh := 0
    q := [][]int{}    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 2 {
                q = append(q, []int{i,j})
            } else if grid[i][j] == 1 {
                fresh++
            }
        }
    }
    if fresh == 0{return 0}
    if len(q) == 0{return -1}
    level := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cr := dq[0]
            cc := dq[1]
            for _, dir := range dirs {
                r := cr + dir[0]
                c := cc + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 1 {
                    q = append(q,[]int{r,c})
                    fresh--
                    grid[r][c] = 2
                }
            }
            qSize--
        }
        level++
    }
    if fresh != 0 {return -1}
    return level-1
}