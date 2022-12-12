func maxDistance(grid [][]int) int {
    
    q := [][]int{}
    max := 0
    m := len(grid)
    n := len(grid[0])
    
    // enqueue all 1's and mark them negative to distinguish between distance 1 and univisted 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                grid[i][j] = -1
                q = append(q, []int{i,j})
            }
        }
    }
    
    // classic bfs processing of queue
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr := dq[0]
        cc := dq[1]
        cv := grid[cr][cc]
        if cv == -1 {
            cv = 0
        }
        
        for _, dir := range dirs {
            r := cr + dir[0]
            c := cc + dir[1]
            if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 0 {
                grid[r][c] = cv + 1
                if cv+1 > max {
                    max = cv+1
                }
                q = append(q, []int{r,c})
            }
        }
    }
    if max == 0 {
        max = -1
    }
    return max
}